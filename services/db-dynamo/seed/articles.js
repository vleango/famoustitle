var params = {
    TableName: 'tech_writer_articles',
    Item: { // a map of attribute name to AttributeValue
        // attribute_value (string | number | boolean | null | Binary | DynamoDBSet | Array | Object)
        // more attributes...
        id: "123",
        author: "dhh",
        title: "Hunting for great names in programming",
        created_at: "2018-04-19T14:49:15.983056333Z",
        tags: ["ruby", "rails", "naming", "web"],
        body: `

One of the real delights of programming is picking great variable, method, and class names. But an even greater treat is when you can name pairs, or even whole narratives, that fit just right. And the very best of those is when you’re forced to trade off multiple forces pulling in different directions. This is the story of one such case.

It started with a simple refactoring. We allow people to post arbitrary URLs in [Basecamp 3](https://basecamp.com/?source=svn)’s chat, which we’ll attempt to fetch and display inline, if its an image, movie, or a Twitter URL. There’s a security concern where we don’t want to allow internal IPs in those URLs, like 127.0.0.1, and then have our Downloader class attempt to trigger an internal request that may sidestep other security precautions.

The specific policy isn’t as important as the fact that this precondition was originally just part of the _Downloader_ class, but now I also needed it in our forthcoming webhooks API. Because just like with previewable chat URLs, webhooks allow users to set URLs that our system then calls. Same underlying security issue to deal with.

No problemo: Just extract the protection into its own module and call it from both spots. First attempt at naming this new module gave me _PrivateNetworkProtector_, which seemed like a fine choice until I considered the method names that’d work with it:

- _PrivateNetworkProtector.protect\_against\_internal\_ip\_address(ip)_
- _PrivateNetworkProtector.verify\_ip\_address\_isnt\_private(ip)_

Hmm. I didn’t like any of those choices. Both are a little too wordy and both include a negative. And come to think of it, I wasn’t even that thrilled with the word _Protector_. It implies something like surge protection, where it just negates the effects of an outlier input. That’s not really what’s going on here. We’re on the lookout for malicious attempts, so a better word would be more forceful. More focused on a threat, not just a risk.

Let’s see what the programmer’s best friend, the thesaurus, had to offer:

![](https://cdn-images-1.medium.com/max/800/1*FUboZOCmwK7__CuYBh6aHA.png)

Lots of possibilities there, but the one that really spoke to me was _Guard_ as in _PrivateNetworkGuard_. Nice. Now I can imagine some big burly fellow checking credentials with an attitude. Just the right image. So back to the right method name. Let’s try the two choices similar to what we had earlier:

- _PrivateNetworkGuard.guard\_against\_internal\_ip\_address(url)_
- _PrivateNetworkGuard.verify\_ip\_address\_isnt\_private(url)_

Hmm, neither of those are right either. I mean, you could use _guard_ again, but the double confetti of this repetition just annoyed my sensibilities. What if we thought of the _Guard_ as something like a prefilter, like _before\_action_ in Action Controller? That’s promising:

- _PrivateNetworkGuard.ensure\_public\_ip\_address(url)_
- _PrivateNetworkGuard.ensure\_no\_private\_ip\_address(url)_

Still not quite right. The _ensure_ verb just doesn’t quite fit with the forceful idea of a _Guard._ It’s too meek. This isn’t a casual checking of credentials. This is IF THIS HAPPENS, I’M RAISING AN EXCEPTION AND BUSTING THE REQUEST!

That lead me to think of the programming language [Eiffel](https://en.wikipedia.org/wiki/Eiffel_%28programming_language%29) and the concept of design by contract. So I browsed the Wikipedia entry to see if I could mine it for a good word. Eiffel uses _require_ to state preconditions, which is what we’re doing here, but that also didn’t see right: _PrivateNetworkGuard.require\_public\_ip_. That’s more like something you’d write in a specification to tell well-meaning actors what they’re supposed to do. This wasn’t about well-meaning type, but rather the nefarious kind.

So I tried the thesaurus again, going off _ensure_ for alternatives:

![](https://cdn-images-1.medium.com/max/800/1*V5sx1v3DkoI2qH0yXDGdhA.png)

No great candidates here. I mean, they’d all work, but they don’t feel quite right. And that’s really what this whole expedition is about. Not just finding something that _could_ work, but something where you go: _Yes! That’s perfect!_ How about something in the realm of _Guard_ then?

![](https://cdn-images-1.medium.com/max/800/1*Oit7eXaSCJAozj-uAx1LPw.png)

Nope. That doesn’t do it either. But one word stood out as an idea: _Police._ The _Guard_ is _policing_ the incoming ip addresses to make sure that an internal one doesn’t slip through. That’s interesting. Let’s keep pulling on that thread:

![](https://cdn-images-1.medium.com/max/800/1*RTzoaE1PQuKJdsa8ObdBVw.png)

_ENFORCE!_ That’s it. Let’s try it out:

- _PrivateNetworkGuard#enforce\_public\_ip(url)_
- _PrivateNetworkGuard#enforce\_no\_private\_ip(url)_

In some ways I like the negative version better, since it strikes to the heart of the responsibility: This is about catching those malicious URLs with private IPs more than its about vetting that something is public. Two sides of the same coin, perhaps, but still an important distinction.

In the end, though, I went with _PrivateNetworkGuard#enforce\_public\_ip(url)_ because I liked keeping a positive method name more than the slightly more apt negative version. Weighing those two subtle trade offs and picking the concern that mattered more.

Now this might seem like a lot of effort to expend searching for a slightly better name, but it goes straight to the heart of programming with a smile. I ventured out to find a great name, not just a passable one. And besides, this whole exercise might have taken five minutes at the most. So not exactly blowing the budget, but definitely lifting my spirit. Isn’t programming great?

* * *
        `
    }
};
docClient.put(params, function(err, data) {
    if (err) ppJson(err); // an error occurred
    else ppJson(data); // successful response
});

var params = {
    TableName: 'tech_writer_articles',
    Item: { // a map of attribute name to AttributeValue
        // attribute_value (string | number | boolean | null | Binary | DynamoDBSet | Array | Object)
        // more attributes...
        id: "456",
        author: "someguy",
        title: "The Lowdown On Absolute vs. Relative Positioning",
        created_at: "2018-04-20T14:49:15.983056333Z",
        tags: ["css", "web"],
        body: `

When I was first learning web development, the style side of CSS seemed straightforward and fun, but performing layout feats seemed like a confusing mess. I sort of felt my way around without a solid understanding of how things like positioning and floats worked and as a result it would take hours to perform even simple tasks. If this situation sounds familiar, then this article is for you.

![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-0.jpg)

One of the real revelations that I had early on was when I was finally able to wrap my head around how positioning contexts worked, especially when it came to the difference between absolute and relative positioning. Today we’re going to tackle this subject and make sure you know exactly how and when to apply a specific positioning context to a given element.

## 5 Different Position Values

![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-3.jpg)

Let’s get some complexity out of the way up front. In reality, there are a whopping **five** different possible values for the _position_ property. We’ll largely skip over _inherit_ because it’s pretty self explanatory (simply inherits the value of its parent) and isn’t really supported well in older versions of IE.

The default _position_ value for any element on the page is _static_. Any element with a _static_ positioning context will simply fall where you expect it to in the flow of the document. This of course entirely depends on the structure of your HTML.

Another value that you’ve no doubt seen is _fixed_. This essentially anchors an object to the background so that wherever you place it, there it will stay. We often see this used on sidebar navigation elements. After scrolling down a long page of content, it’s a pain to go back to the top to navigate elsewhere so applying a fixed position to the navigation means the user never loses site of the menu. Click the image below to see a live example of this in action.

[![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-1.jpg)](http://tinkerbin.com/NMUpJJZl)

So there you have three _position_ values that are fairly straightforward and easy to wrap your mind around. The final two are of course _absolute_ and _relative_. Let’s focus on explaining these independently and then take a look at how they can be used together to achieve unique results.

## Absolute Positioning

Absolute positioning allows you to remove an object from the typical flow of the document and place it at a specific point on the page. To see how this works, let’s set up a simple unordered list of items that we’ll turn into clearly discernible rectangles.

![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-2.jpg)

As we’ve already learned, by default these items will have a static position applied. That means they follow the standard flow of the document and are positioned according to the margins that I’ve placed on the list. Now let’s see what happens if I target one of these list items and apply a value of _absolute_ to the position property.

![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-4.jpg)

As you can see, the fourth list item was completely removed from the normal flow and placed at the top left corner of the browser window. Note that even if there were other content occupying this position, this element wouldn’t care. When something has absolute positioning applied, it neither affects nor is affected by other elements in the normal flow of the page.

The reason for absolute positioning is so we can position this item precisely where we want it. We do this with the top, bottom, left and right CSS properties. For instance, let’s say we wanted the fourth list item to be placed twenty pixels from the topside of the browser window and twenty pixels from the left side.

![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-5.jpg)

To prove our earlier point about absolutely positioned items not interacting with other content, let’s move the fourth list item right into the territory of the other list items. Watch how it simply overlaps the existing content instead of pushing it around. Click on the image below to see and tweak a live example of this test.

[![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-6.jpg)](http://tinkerbin.com/GqfhucDo)

As one final note, notice how the fifth list item occupies the position previously held by the fourth rather than holding its position as if the fourth were still in place. Since the fourth item has been removed from the flow, everything else will adjust accordingly.

## Relative Positioning

Relative positioning works similarly to absolute positioning in that you can use top, bottom, left and right to scoot an object to a specific point on the page. The primary difference is the origin or starting point for the element. As we saw above, with absolute positioning, the starting point was at the very top left of the browser window. Check out what happens when we apply relative positioning though:

![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-7.jpg)

Nothing happened! Or did it? As it turns out, the object is indeed relatively positioned now, but its starting point is where it normally lies in the flow of the document, not the top left of the page. Now if we apply the same 20px adjustments that we used before, the result is quite different.

![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-8.jpg)

This time the item was moved “relative” to its starting position. This is extremely helpful for when you want to slightly tweak an object’s position rather than completely reset it altogether. Notice that, just as with absolute positioning, the relatively positioned object doesn’t care about other items in the normal low of the page. However, the original position occupied by the relatively positioned element hasn’t been occupied by the next list item as it did with absolutely positioned elements, instead the document acts as if the fourth item still occupies that position.

## Putting Them Together

Now that you know how absolute and relative positioning work on their own, it’s time to dive into a more complex example and see how they can work together in a truly useful manner. This time we’re going to build a nice little demo to show off how it all works.

### HTML

Start with a simple div with a class of “photo” and place a 200x200px image inside. This is all the HTML we need so after you’ve got this, move over to some CSS.

1

### Basic CSS

In your CSS, start by changing the body color to something dark. Then apply some basic styles to the photo element and give it some nice border and shadow styles.

1

Here’s the resulting image. It’s nothing special, but it will provide a great testing ground for our positioning experiment.

![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-9.jpg)

### Applying a Strip of Tape

Let’s say we wanted to create the illusion that this photo was hanging from the background, connected by a small strip of tape. To pull this off, we’ll need to flex our newfound positioning muscle and leverage some pseudo elements.

The first thing we want to do is use the :before pseudo element to create our strip of tape. We’ll give it a height of 20px and a width of 100px, then set the background to white at 50% opacity. I’ll finish by adding in a slight box-shadow.

  .photo:before {
    content: "";
    height: 20px;
    width: 100px;
    background: rgba(255,255,255,0.5);

    -webkit-box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
       -moz-box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
            box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
  }

&nbsp;

If we look at a live preview of our page after this code, we can see that we really screwed up our image. The piece of tape is really interfering with the flow of the document. Even though it’s not really visible, it has bumped our image right out of its border!

![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-10.jpg)

Obviously, we’ve got some issues with how the pseudo element is being positioned. To attempt to fix this, let’s see what happens if we apply relative positioning to the piece of tape.

  .photo:before {
    content: "";
    height: 20px;
    width: 100px;
    background: rgba(255,255,255,0.5);
    position: relative;
    top: 0px;
    left: 0px;

    -webkit-box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
       -moz-box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
            box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
  }

&nbsp;

Here’s the effect of this code:

![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-10.jpg)

As you can see, we didn’t fix our problem, everything is just as screwed up as before. Since this didn’t work, let’s take the opposite route and see what happens if we use absolute positioning.

  .photo:before {
    content: "";
    height: 20px;
    width: 100px;
    background: rgba(255,255,255,0.5);
    position: absolute;
    top: 0px;
    left: 0px;

    -webkit-box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
       -moz-box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
            box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
  }

&nbsp;

Here’s what our demo looks like now:

![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-11.jpg)

Our tape has finally made an appearance! Unfortunately, it’s way up in the corner. We could nudge it into place with the top and left values, but that wouldn’t actually be a workable solution. The reason is that this would put the tape at a specific point on the background, where it would stay no matter what. However, the image has been automatically centered in the viewport so as you change the window size it actually moves, meaning the piece of tape won’t be correctly positioned anymore.

So now we’ve tried both relative and absolute positioning with neither providing the solution we want. Where do we turn next? As it turns out, we haven’t really gotten the full story behind absolute positioning yet. You see, it doesn’t always default to the top left of the browser window. Instead, what _position: absolute;_ really does is position the element **relative to its first non-statically-positioned ancestor** (inherit doesn’t count either). Since there hasn’t been one of those in previous examples, it was simply reset to the origin of the page.

So how does this translate into useful information? It turns out, we can use absolute positioning on our piece of tape, but we first need to add a positioning context to its ancestor, the photo. Now, we don’t want to absolutely position that element because we don’t want it to move anywhere. Thus, we apply relative positioning to it. Let’s see what this looks like.

  .photo {
    margin: 50px auto;
    border: 5px solid white;
    width: 200px;
    height: 200px;
    position: relative;

    /*overly complex but cool shadow*/
    -webkit-box-shadow: 0px 10px 7px rgba(0,0,0,0.4), 0px 20px 10px rgba(0,0,0,0.2);
       -moz-box-shadow: 0px 10px 7px rgba(0,0,0,0.4), 0px 20px 10px rgba(0,0,0,0.2);
            box-shadow: 0px 10px 7px rgba(0,0,0,0.4), 0px 20px 10px rgba(0,0,0,0.2);
  }


  .photo:before {
    content: "";
    height: 20px;
    width: 100px;
    background: rgba(255,255,255,0.5);
    position: absolute;
    top: 0px;
    left: 0px;

    -webkit-box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
       -moz-box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
            box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
  }

&nbsp;

As you can see, the pseudo element has absolute positioning applied, which means it will choose a starting point given the location of its first non-static ancestor. Since I’ve applied relative positioning to the photo, that item fits this description. So now our piece of tape will begin at the origin of the photo (even if the photo moves due to browser resizing).

![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-12.jpg)

Now that we have found the starting position that we were looking for, we can tweak the top and left values to nudge the tape into place. Notice that I’ve applied a negative value to the top property so the tape sticks out of the image onto the background. The left position is set to center the tape horizontally ([click here](http://new.instacalc.com/1342) to see how the math works out).

  .photo:before {
    content: "";
    height: 20px;
    width: 100px;
    background: rgba(255,255,255,0.5);
    position: absolute;
    top: -15px;
    left: 50px;

    -webkit-box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
       -moz-box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
            box-shadow: 0px 1px 3px rgba(0,0,0,0.4);
  }

&nbsp;

As we can see in the finished version below, combining absolute and relative positioning was exactly what we needed to pull off the effect that we were going for. Click the image below to see and tweak the live demo.

[![screenshot](https://cmv-ds-images.s3.amazonaws.com/wp-content/uploads/fixedvsrelative-13.jpg)](http://tinkerbin.com/bynpIxUx)

## A Quick Summary

Positioning contexts can be difficult to wield if you’re attempting to implement them without a solid foundation of knowledge behind how they work. Fortunately, there are only three primary pieces of information that you need to remember to master absolute and relative positioning.

The first is that relative positioning will allow you to tweak an element’s position relative to its normal starting point. The second is that absolute positioning will allow you to tweak an element’s position relative to its first non-statically-positioned ancestor (defaults to page bounds if none is found). The final piece of information to remember is that both relatively and absolutely positioned items won’t affect the static and fixed items around them (absolutely positioned items are removed from the flow, relatively positioned items occupy their original position).
        `
    }
};
docClient.put(params, function(err, data) {
    if (err) ppJson(err); // an error occurred
    else ppJson(data); // successful response
});
