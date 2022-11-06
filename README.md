# Cache
This little library knows how to store values ​​in a map

example: 
    c:= cache.New()
    err = c.Set("ID",1)
    value, err = c.Get("ID")
    err = c.Delete("ID")

    if err != nil{
     fmt.Println(err.Error())   
     fmt.Println(err.Unwrap())
    }