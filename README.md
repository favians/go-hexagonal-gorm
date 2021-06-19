# Sample Golang API Server

Sample REST API build using echo server.

The code implementation was inspired by port and adapter pattern or known as [hexagonal](blog.octo.com/en/hexagonal-architecture-three-principles-and-an-implementation-example):

-   **Business**<br/>Contains all the logic in domain business. Also called this as a service. All the interface of repository needed and the implementation of the service itself will be put here.
-   **Modules**<br/>Contains implementation of interfaces that defined at the business (also called as server-side adapters in hexagonal's term)
-   **API**<br/>API http handler or controller (also called user-side adapters in hexagonal's term)

# Data initialization

To describe about how port and adapter interaction (separation concerned), this example will have two databases supported. There are MySQL and MongoDB.

MongoDB will become a default databaese in this example. If you want to change into MySQL, update the configuration inside
[config.yaml](https://raw.githubusercontent.com/muhsinshodiq/golang-sample-api/master/config/config.yaml) file.

### MongoDB

Please execute script below to create a new collection called `items` including the index needed

```mongodb
db.createCollection('items');
db.items.createIndex({"tags": 1});
db.items.createIndex({"modified_at": 1, "_id": 1});
```

### MySQL

Please execute script below to create `item` and `item_tag` table in your database

```sql
CREATE TABLE `item` (
  `id` varchar(24) NOT NULL DEFAULT '',
  `name` text NOT NULL,
  `description` text NOT NULL,
  `created_at` datetime NOT NULL,
  `created_by` varchar(50) NOT NULL DEFAULT '',
  `modified_at` datetime NOT NULL,
  `modified_by` varchar(50) NOT NULL DEFAULT '',
  `version` int(11) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `modified_at` (`modified_at`,`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `item_tag` (
  `item_id` varchar(24) NOT NULL DEFAULT '',
  `tag` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`item_id`,`tag`),
  KEY `tag` (`tag`),
  CONSTRAINT `item_tag_ibfk_1` FOREIGN KEY (`item_id`) REFERENCES `item` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
```

# How To Run Server

Just execute code below in your console

```console
./run.sh
```

# How To Consume The API

There are 4 availables API that ready to use:

-   GET `/v1/items/:id`
-   GET `/v1/items/[tag-name]`
-   POST `/v1/items`
-   PUT `/v1/items`

To make it easier please download [Insomnia Core](https://insomnia.rest) app and import [this collection](https://raw.githubusercontent.com/muhsinshodiq/golang-sample-api/master/insomnia.json).
