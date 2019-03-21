package model

import (
    "github.com/firedial/midas-go/db"
)

type Attribute struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    Group_id int `json:"group_id"`
}

func GetAttribute(attribute_name string, id string) []Attribute {
    var attributes []Attribute

    db := db.Init();
    defer db.Close();

    where := "";
    if id != "" {
        where = " WHERE id = " + string(id)
    }

    _ = where

    stmt, err := db.Prepare("SELECT id, name, description, group_id FROM " + attribute_name)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    rows, err := stmt.Query()
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var attribute Attribute

        err := rows.Scan(
            &(attribute.Id),
            &(attribute.Name),
            &(attribute.Description),
            &(attribute.Group_id))
        if err != nil {
            panic(err)
        }

        attributes = append(attributes, attribute)
    }

    if err != nil {
        panic(err.Error())
    }
    return attributes
}


