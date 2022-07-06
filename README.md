# What is Leago?
**Leago** is a small and simple library with a convenient interface to interact with Riot API. 

![League of Legends picture](https://coop-land.ru/uploads/posts/2020-01/1578739860_1.jpeg)

## Notes
* Riot API has many more methods that this implementation written in Goland.
* Before using the library, it's advised to visit Riot API documentation page.

# How to start using?
<pre>go get github.com/mikemight/leago</pre>

# Finally
Leago uses a rate limiter from [this library].(https://github.com/throttled/throttled/)

**IMPORTANT:** When initializing the API client, use restrictions that match the Riot API restrictions!
* 20 requests every 1 seconds(s)
* 100 requests every 2 minutes(s)

Otherwise, very often there will be an error associated with exceeding the number of requests.

# Riot Developer Portal | RDP
[Go to view](https://developer.riotgames.com/docs/portal)
