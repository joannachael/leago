# What is Leago?
Leago is a small library with a simple interface to interact with the Riot API.
The important thing is that the Riot API has many more methods than this implementation written in Golang.

**Why is that?** This library is intended to use a part of the API, the functionality of which is to receive data from League of Legends.

**Note:** Before using library, it is advised to visit the Riot API documentation page to have a cursory understanding of the data it returns.
![LoL](https://coop-land.ru/uploads/posts/2020-01/1578739860_1.jpeg)


# How to start using?
<pre>go get github.com/mikemight/leago</pre>
After reading the documentation, the data stored in the "regions" and "leagues" files will improve your understanding of which values are used to query.

# Finally
The Leago library uses the rate limiter from [here.](https://github.com/throttled/throttled/)

**IMPORTANT:** When initializing the API client, use restrictions that match the Riot API restrictions!
* 20 requests every 1 seconds(s)
* 100 requests every 2 minutes(s)

Otherwise, you will often be greeted with an error when making a request.

# Riot Developer Portal | RDP
[Go to view](https://developer.riotgames.com/docs/portal)
