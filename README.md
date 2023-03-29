# groupie-tracker-geolocalization

### Description

Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.

- It was given an [API](https://groupietrackers.herokuapp.com/api), that consists in four parts:

  - The first one, `artists`, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

  - The second one, `locations`, consists in their last and/or upcoming concert locations.

  - The third one, `dates`, consists in their last and/or upcoming concert dates.

  - And the last one, `relation`, does the link between all the other parts, `artists`, `dates` and `locations`.

Given all this a user friendly website was built where you can see the bands info through several data visualizations.

The **filter** lets the user filter the artists/bands that will be shown. Four filters are incorporated:
  > number of members(1 to 8), locations of concerts, first album date,  and creation date.


Groupie Tracker Geolocalization consists on mapping the different concerts locations of a certain artist/band given by the Client.
  > A process of converting addresses (ex: Germany Mainz) into geographic coordinates (ex: 49,59380 8,15052) is used to place markers for the concerts locations of a certain artist/band on a map.

### Instructions

- To run the program write: `$go run cmd/main.go`

- Open `http://localhost:8080`


### This project implements :

- Manipulation, display and storage of data.
- Manipulation of Maps API
- [JSON](https://www.json.org/json-en.html) files and format.
- HTML, CSS
- [Client-server](https://developer.mozilla.org/en-US/docs/Learn/Server-side/First_steps/Client-Server_overview).
- Geolocation, geocoding, etc


### 💎 Authors
 - ashagiro and MeirbekAshirbayev