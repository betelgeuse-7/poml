## POML (Paranthesis-Obsessed Markup Language)

A markup language that has a one-to-one relationship with HTML.

```lisp
    (p "Hello")
```
is
```html
<p>Hello</p>
```
---

```lisp
    (tag-name [:attr "attr-val"] [child-elements])
```

##### Examples
```lisp
(div 
    ; this is a comment
    (h3 "A \"Cat\" Picture")
    (a :href "https://google.com" "Google")
    (div :class "cat-div container"
        (img :id "catphoto" :src "https://example.com/img/cat.jpg")
    )
    (button :onclick "doSomething()" "Click Me" :style "background-color: blue; border-radius: 3px;")
)
```
```html
<div>
    <!-- this is a comment-->
    <h3>A "Cat" Picture</h3>
    <a href="https://google.com">Google</a>
    <div class="cat-div container">
        <img id="catphoto" src="https://example.com/img/cat.jpg">
    </div>
    <button onclick="doSomething()" style="background-color: blue; border-radius: 3px;">Click Me</button>
</div>
```