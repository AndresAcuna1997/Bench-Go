Las interfaces son un type especial el cual se asegura que el el tipo que estamos  pasando posea un método en especifico

```go
type saver interface {
  Save() error
}

func saveData(data saver) error {
  err := data.Save()

  if err != nil {
    fmt.Println("Error saving the note")
    return err
  }

  fmt.Println("File saved!")
  return nil
}
```

En el caso anterior la función `saveData()` acepta un argumento de tipo saver lo cual significa.  El argumento `data` sin importar de que type sea Go va a revisar que este tenga un método llamado `Save` 