<h1>A simple in memory cache implementation</h1>
The simple implementation of a in memory cache that uses a slice as a data structure that stores data as {key string} value interface{}
It implements such methods as:
<ul>
<li><b>New</b> - to create an instance of a Cache structure</li>
<li><b>Set</b> - to add a new value by key or replace existing. The operation is done in O(n)</li>
<li><b>Delete</b> - to delete a value by key if it exists. The operation is done in O(1)</li>
<li><b>Size</b> - to get a number of elements stored in a cache at the moment</li>
</ul>
<h1>How to install</h1>

```sh
go get github.com/sasha-filippov/cache
```

<h1>How to use</h1>
to initialize 

```sh
cache := cache.New()
```

to add a new element:

```sh
cache.Set("someKey", anyValue)
```

to check how many elements:

```sh
fmt.Println(cache.Size())
```

to delete an element:

```sh
cache.Delete("somekey")
```
