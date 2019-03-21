package controller

import(
    "github.com/firedial/go-midas/model"
)

func AttributeGet(attribute_name string, queries map[string][]string) model.Attribute {
    id := queries["id"][0]
    attribute := model.GetAttribute(attribute_name, id)
    return attribute
}