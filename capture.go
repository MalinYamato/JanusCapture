
package main

import (
	"github.com/jmoiron/jsonq"
	"net/http"
	"encoding/json"
	"bytes"

	"fmt"

	"strconv"
	"io/ioutil"
	"strings"

)

type JanusRequest struct {
	Janus string        `json:"janus"`
	Transation string   `json:"transaction"`
	Secret     string   `json:"admin_secret"`

}
type JanusSessions struct {
	Janus      string   `json:"janus"`
	Transation string   `json:"transaction"`
	Sessions   []int    `json:"sessions"`

}
type JanusHandles struct {
	Janus      string   `json:"janus"`
	Transation string   `json:"transaction"`
	Session    int      `json:"session"`
	Handles    []int    `json:"handles"`
}

func getDocument(url string, message JanusRequest) (r *http.Response) {
	fmt.Println(url)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(message)
	res, _ := http.Post(url, "application/json; charset=utf-8", b)
	return res
}

type handleID int
type privateID int

type Publishment struct
{
	RoomID    int
//	Audio     bool  // NOT in use
//  Video     bool  // NOT in use
}
type Subscription struct {
	RoomID    int
	ID        int        //subscriber
	Display   string     //display of subscriber
	HandleID  handleID
	PrivateID privateID  //owner of feed
//	Audio     bool  // NOT in use
//	Video     bool  // NOT in use
}
type MediaUser struct {
	 ID            int
	 PrivateID     privateID
	 Display       string
	 SessionID     int
	 Publishments  map[handleID]Publishment
	 Subscriptions map[handleID]Subscription
}
type MediaUsers struct {
	__mus map[string]MediaUser
}
func (mus *MediaUsers) findByDisplay(display string) (MediaUser, bool) {
	mu, err := mus.__mus[display]
	return mu, err
}
func (mus *MediaUsers) update(mu MediaUser) {
	mus.__mus[mu.Display] = mu
}
func (mus *MediaUsers) listeners(mu MediaUser) ([]MediaUser) {
	result := []MediaUser{}
	for _, mediaUser := range mus.__mus {
		for _, aSubby := range mediaUser.Subscriptions {
			if aSubby.PrivateID == mu.PrivateID {
				result = append(result, mediaUser)
			}
		}
	}
	return result
}


func main () {

	mediaUsers := MediaUsers{ map[string]MediaUser{}}
	subscriptions := map[handleID]Subscription{}

	url := "https://media.raku.cloud:7889/admin"
	message := JanusRequest{Janus: "list_sessions", Transation: "123", Secret: "janusoverlord"}
	var res= getDocument(url, message)
	var sessions JanusSessions
	err := json.NewDecoder(res.Body).Decode(&sessions)
	if err != nil {
		fmt.Println("err")
	}
	for i := 0; i < len(sessions.Sessions); i++ {
		var handles JanusHandles
		message.Janus = "list_handles"
		res = getDocument(url+"/"+strconv.Itoa(sessions.Sessions[i]), message)
		err := json.NewDecoder(res.Body).Decode(&handles)
		if err != nil {
			fmt.Println("err")
		}
		for h := 0; h < len(handles.Handles); h++ {
			message.Janus = "handle_info"
			res = getDocument(url+"/"+strconv.Itoa(sessions.Sessions[i])+"/"+strconv.Itoa(handles.Handles[h]), message)
			body, _ := ioutil.ReadAll(res.Body)
			//fmt.Println(string(body))
			data := map[string]interface{}{}
			dec := json.NewDecoder(strings.NewReader(string(body)))
			dec.Decode(&data)
			jq := jsonq.NewQuery(data)
			//fmt.Println(string(body))
			pubsub, _ := jq.String("info", "plugin_specific", "type")
			if (pubsub == "publisher") {
				id, _ := jq.Int("info", "plugin_specific", "private_id")
				private_id := privateID(id)
				display, _ := jq.String("info", "plugin_specific", "display")
				var aPublisher MediaUser
				_, ok := mediaUsers.findByDisplay(display)
				if ! ok {
					aPublisher := MediaUser{}
					aPublisher.Publishments = map[handleID]Publishment{}
					aPublisher.SessionID, _ = jq.Int("session_id")
					aPublisher.Display, _ = jq.String("info", "plugin_specific", "display")
					aPublisher.ID, _ = jq.Int("info", "plugin_specific", "id")
					aPublisher.PrivateID = private_id
					mediaUsers.update(aPublisher)
				}
				aPublisher, _ = mediaUsers.findByDisplay(display)
				_, err := jq.Int("info", "streams", "0", "id")
				if err != nil {
					fmt.Println("no publishing streams")
				} else {
					// the publisher is broadcasting
					room, _ := jq.Int("info", "plugin_specific", "room")
					id, _ = jq.Int("handle_id")
					handle_id := handleID(id)
					aPublisher.Publishments[handle_id] = Publishment{room}
				}
				mediaUsers.update(aPublisher)

			} else if (pubsub == "listener") {
				id, _ := jq.Int("handle_id")
				handle_id := handleID(id)
				id, _ = jq.Int("info", "plugin_specific", "private_id")
				private_id := privateID(id)

				_, err := jq.Int("info", "streams", "0", "id")
				if err != nil {
					// the listener is not listening
					fmt.Println("no listening streams")
				} else {
					_, ok := subscriptions[handle_id]
					if ! ok {
						subscriptions[handle_id] = Subscription{}
					}

					room, _ := jq.Int("info", "plugin_specific", "room")
					id, _ = jq.Int("handle_id")
					handle_id := handleID(id)
					subby := subscriptions[handle_id]
					subby.RoomID = room
					subby.PrivateID = private_id
					subby.ID, _ = jq.Int("info", "plugin_specific", "feed_id")
					subby.Display, _ = jq.String("info", "plugin_specific", "feed_display")
					subby.HandleID = handle_id
					subscriptions[handle_id] = subby
				}
			}
		}
	}


	for _, user := range mediaUsers.__mus {
		for _, subby := range subscriptions {
			if user.PrivateID == subby.PrivateID {
				if (user.Subscriptions == nil) {
					user.Subscriptions = map[handleID]Subscription{}
				}
				user.Subscriptions[subby.HandleID] = subby
				mediaUsers.update(user)
			}
		}
	}

	for _, user := range mediaUsers.__mus {
		fmt.Print("User: ")
		fmt.Printf("Display %s ID %d PvtID %d  Session %d\n",user.Display, user.ID, user.PrivateID, user.SessionID )
		fmt.Println("publishes: ")
		for h, pub := range user.Publishments {
			fmt.Printf("Using handle %d in Room %d \n", h, pub.RoomID)
		}
		fmt.Println("subscribes to: ")
		for s, sub := range user.Subscriptions {
			fmt.Printf("Using handle %d in  Room %d to %s with ID %d PvtID %d\n", s, sub.RoomID, sub.Display, sub.ID, sub.PrivateID)
		}
      fmt.Println()
	}

//	fmt.Println(mediaUsers)
//	fmt.Println(subscriptions)

}