package main

import (
  "net/http"
)

// Client does client stuff.
type ProviderClient struct {
  endpoint string
}

func createRequest(id string, config interface{}) error {
  // TODO POST to api with payload
  return nil
}

func createDeploy(id string, config interface{}) error {
  // TODO POST to api with payload
  return nil
}

func getRequest(id string, meta interface{}) error {
  // TODO GET to api, parse response into request/deploy objects and return
  return nil
}

func waitForRequest(id string, status string, meta interface{}) error {

  activeState := false

  // TODO actually parse response and compare status

  // var endpoint = meta.(*Conf).endpoint + "/requests/request/" + id

  for activeState == false {
    // res, err := http.Get(endpoint)

    // if err != nil {
    //  return err
    // }

    // defer r.Body.Close()
    // decoder := json.NewDecoder(res.Body)

    // activeState = *res.Table.TableStatus == "ACTIVE"

    // // Wait for a few seconds
    // if !activeState {
    //  log.Printf("[DEBUG] Sleeping for 5 seconds for table to become active")
    //  time.Sleep(5 * time.Second)
    // }
  }

  return nil
}

func deleteRequest(id string, meta interface{}) error {
  var endpoint = meta.(*ProviderClient).endpoint + "/requests/request/" + id
  req, _ := http.NewRequest("DELETE", endpoint, nil)
  _, err := http.DefaultClient.Do(req)
  // TODO wait until done
  return err
}


