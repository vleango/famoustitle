export const articles = [
  {
    id: '123',
    title: 'Hunting for great names in programming',
    author: 'dhh',
    body: `One of the real delights of programming is picking great variable, method, and class names. But an even greater treat is when you can name pairs, or even whole narratives, that fit just right. And the very best of those is when you’re forced to trade off multiple forces pulling in different directions. This is the story of one such case.

It started with a simple refactoring. We allow people to post arbitrary URLs in Basecamp 3’s chat, which we’ll attempt to fetch and display inline, if its an image, movie, or a Twitter URL. There’s a security concern where we don’t want to allow internal IPs in those URLs, like 127.0.0.1, and then have our Downloader class attempt to trigger an internal request that may sidestep other security precautions.

The specific policy isn’t as important as the fact that this precondition was originally just part of the Downloader class, but now I also needed it in our forthcoming webhooks API. Because just like with previewable chat URLs, webhooks allow users to set URLs that our system then calls. Same underlying security issue to deal with.

No problemo: Just extract the protection into its own module and call it from both spots. First attempt at naming this new module gave me PrivateNetworkProtector, which seemed like a fine choice until I considered the method names that’d work with it:

`,
    created_at: '2018-02-28T14:48:54.444740278Z',
    tags: ["ruby", "rails", "naming", "web"]
  },
  {
    id: '456',
    title: 'Creating staging and other environments in Rails',
    author: 'Josef Strzibny',
    body: `
    Ruby on Rails come with three environments by default – development, testing and production. But sooner or later one has a need for staging environment. And don’t get me wrong, you can (or should?) use the production settings there, but if you run it locally or on the same server as production, chances are you need a different database. And while you are at it, it may be handy to allow logging to console or change any other of Rails settings for that matter. In fact you can create as many other environments as you want and since it’s really easy I encourage you to do so.

To create a new environment you need to create:

a new config/environments/YOUR_ENVIRONMENT.rb file
a new database configuration entry in config/database.yml if your application uses database
a new secret key base entry in config/secrets.yml for apps on Rails 4.1 and higher
As I mentioned first we would need a new file in config/environments/. A short example for staging environment could be:
    `,
    created_at: '2018-02-28T14:48:54.444740278Z',
    tags: ["css", "web"]
  }
]

export const article = articles[0]
