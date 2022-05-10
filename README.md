## POML (Paranthesis-Obsessed Markup Language)

A markup language that has a one-to-one relationship with HTML.

```lisp
    (p text: Hello)
```
is
```html
<p>Hello</p>
```
---

```lisp
    (tag-name [attr:attr-val;] [child-elements])
```

##### Examples
```sml
(div 
    (* this is a comment)
    (h3 text: A Cat Picture;)
    (a href: https://google.com; text: Google;)
    (div class: cat-div container;
        (img id: catphoto; src: https://example.com/img/cat.jpg;)
    )
    (button onclick: doSomething(); text: Click Me;)
)
```
```html
<div>
    <h3>A Cat Picture</h3>
    <a href="https://google.com">Google</a>
    <div class="cat-div container">
        <img id="catphoto" src="https://example.com/img/cat.jpg">
    </div>
    <button onclick="doSomething()">Click Me</button>
</div>
```