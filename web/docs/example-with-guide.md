# Example with guide

This guide will show minimal setup with page, multiple component instances, lifecycle integration and `net/http` setup. This guide will rely on demo project setup, that can be found [here](https://github.com/yuriizinets/ssceng/tree/master/demo)

## Entry point

First, we need to setup serving basis

```go
func main() {
    // Init serve mux
    mux := http.NewServeMux()

    // Routes
    // ...

    // Run
    if os.Getenv("PORT") == "" {
        log.Println("Listening on localhost:25025")
        http.ListenAndServe("localhost:25025", mux)
    } else {
        log.Println("Listening on 0.0.0.0:" + os.Getenv("PORT"))
        http.ListenAndServe(":"+os.Getenv("PORT"), mux)
    }
}
```

## Page handler

After that, we need to define generic page handler.  
SSC is not providing own page handler, because you can use any serving basis, not only `net/http`.
Instead of that, SSC provides PageHandlerFactory, helper for creating own page handler.

```go
func pagehandler(p ssc.Page) http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        ssc.PageHandlerFactory(p, map[string]interface{}{
            "internal:rw": rw,
            "internal:r":  r,
        })(rw, r)
    }
}
```

## Page

Now, we can define our page.  
Page is represented by a struct which implements [Page](https://github.com/yuriizinets/ssceng/blob/master/types.go#L51) interface.
Page requires method, returning ready for use template. In this example, we will store our page markup in `page.index.html`.
`ssc.Funcs` is a function, that returns FuncMap. This funcmap is required for correct work of some SSC features.

```go
type PageIndex struct {}

func (p *PageIndex) Template() *template.Template {
    return template.Must(template.New("page.index.html").Funcs(ssc.Funcs()).ParseGlob("*.html"))
}
```

page.index.html

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>SSC Quick Start</title>
</head>

<body>
    ...
</body>
</html>
```

## Component

Let's define sample component, which fetches UUID from httpbin page.  
Component is represented by a struct which implements [Component](https://github.com/yuriizinets/ssceng/blob/master/types.go#L57) interface.
By default, Component interface doesn't have any required methods. Instead of having all-in-one, we have multiple interfaces with functionality separation.
This approach also covers pages. In this example, we will implement [ImplementsAsync](https://github.com/yuriizinets/ssceng/blob/master/types.go#L69) interface.
This method will be called as goroutine in page rendering [lifecycle](/concepts.html#lifecycle).
In that way, all needed async data will be fetched concurrently. In this example, component's markup will be stored in `component.uuid.html`

```go
type ComponentUUID struct {
    UUID string
}

func (c *ComponentUUID) Async() error {
    resp, err := http.Get("http://httpbin.org/uuid")
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    data := map[string]string{}
    json.NewDecoder(resp.Body).Decode(&data)
    c.UUID = data["uuid"]
    return nil
}
```

component.uuid.html

```html
{{ define "ComponentUUID" }}
<div>
    httpbin.org uuid: {{ .UUID }}
</div>
{{ end }}
```

## Attaching component

For using component, you need to define page fields for storing component objects and `Init` method for initialization and registration of components.
Inside of init, use `ssc.RegC` for registering your components. In that way you're including component in page render [lifecycle](/concepts.html#lifecycle).
After that, you need to pass component object to template in your page markup.

```go
type PageIndex struct {
    DemoUUID1 ssc.Component
    DemoUUID2 ssc.Component
    DemoUUID3 ssc.Component
    DemoUUID4 ssc.Component
}

...

func (p *PageIndex) Init() {
    p.DemoUUID1 = ssc.RegC(p, &ComponentUUID{})
    p.DemoUUID2 = ssc.RegC(p, &ComponentUUID{})
    p.DemoUUID3 = ssc.RegC(p, &ComponentUUID{})
    p.DemoUUID4 = ssc.RegC(p, &ComponentUUID{})
}
```

page.index.html

```html
...
<body>
    {{ template "ComponentUUID" .DemoUUID1 }}
    {{ template "ComponentUUID" .DemoUUID2 }}
    {{ template "ComponentUUID" .DemoUUID3 }}
    {{ template "ComponentUUID" .DemoUUID4 }}
</body>
...
```

## Attaching page

For attaching your page, now you can simply use page handler bellow `Routes` comment in your main function.

```go
...
mux.HandleFunc("/", pagehandler(&PageIndex{}))
...
```

## Running

Ready! Now you can run your app with usual `go run .`  
For setting custom port, or exposing on local network, you can run in that way `PORT=25025 go run .`