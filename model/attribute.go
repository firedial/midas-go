package model

import (
    "github.com/firedial/go-midas/db"
)

type Attribute struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    Group_id int `json:"group_id"`
}

func GetAttribute(attribute_name string, id string) Attribute {
    var attribute Attribute

    db := db.Init();
    defer db.Close();

    where := "";
    where = " WHERE id = " + string(id)

    err := db.QueryRow(`
        SELECT
            id,
            name,
            description,
            group_id
        FROM ` + attribute_name + where).Scan(
        &(attribute.Id),
        &(attribute.Name),
        &(attribute.Description),
        &(attribute.Group_id))

  if err != nil {
    panic(err.Error())
  }
  return attribute
}


