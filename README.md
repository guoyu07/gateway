Gateway
--------

`gateway` is a open sourced gateway implementation providing
`authentication`, `throttling`, `load balancing` to backend services.

# Why another gateway?

In Service-Oriented Architecture, a gateway usually works as a load balancer
and service dispatcher, and provides common but trivial functionalities,
such as authentication and throttling, to backend services.

You're able to build a gateway easily with 3rd party utilities or libraris
such as [OpenResty](https://openresty.org/en/) and [Micro](https://github.com/micro/micro).

But when you need to share stateful data
across distributed instances of the gateway, typically,
they need rely on a centralized storage service, such as etcd or redis.

Unfortunately, if you build a gateway in this way,
both flexibility and availability of the gateway relys on the centralized storage.
For example, if you build a gateway upon Redis storage,
the total QPS of your gateway will be limited by capacity of Redis.

The motivation of this project is to build a storage-free `gateway`,
which is highly available and flexible.
And the capacity of `gateway` should be able to be scaled to N times by
simply adding more instances into the cluster.

To achieve this, a decentralized and built-in event layer in `gateway` provides
highly available and efficient communication channel across the distributed system,
so that instances of `gateway` can co-operate with each other in an even-driven
way without relying on any centralized storage systems.

# Features

I'd like to have following features(plugable) in V1:

- [x] generic token bucket throttler

- [x] authenticator

- [x] round robin load balancer for backend services


# Contribute to this project

Submit your issue or merge request, we'll do our best to feedback timely.

