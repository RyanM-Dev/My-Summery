### Name

mqtt — MQ Telemetry Transport

### Synopsis

`MQTT`

### Description

**MQTT** is a lightweight publish/subscribe messaging protocol. It is useful for use with low power sensors, but is applicable to many scenarios.

This manual describes some of the features of MQTT version 3.1.1/3.1, to assist end users in getting the most out of the protocol. For more complete information on MQTT, see `https://mqtt.org/`.

### Publish/Subscribe

The MQTT protocol is based on the principle of publishing messages and subscribing to topics, or "pub/sub". Multiple clients connect to a broker and subscribe to topics that they are interested in. Clients also connect to the broker and publish messages to topics. Many clients may subscribe to the same topics and do with the information as they please. The broker and MQTT act as a simple, common interface for everything to connect to. This means that you if you have clients that dump subscribed messages to a database, to Twitter, or even a simple text file, then it becomes very simple to add new sensors or other data input to a database, Twitter or so on.

### Topics/Subscriptions

Messages in MQTT are published on topics. There is no need to configure a topic, publishing on it is enough. Topics are treated as a hierarchy, using a slash (/) as a separator. This allows sensible arrangement of common themes to be created, much in the same way as a filesystem. For example, multiple computers may all publish their hard drive temperature information on the following topic, with their own computer and hard drive name being replaced as appropriate:

- sensors/COMPUTER_NAME/temperature/HARDDRIVE_NAME
    

Clients can receive messages by creating subscriptions. A subscription may be to an explicit topic, in which case only messages to that topic will be received, or it may include wildcards. Two wildcards are available, `+` or `#`.

`+` can be used as a wildcard for a single level of hierarchy. It could be used with the topic above to get information on all computers and hard drives as follows:

- sensors/+/temperature/+
    

As another example, for a topic of "a/b/c/d", the following example subscriptions will match:

- a/b/c/d
    
- +/b/c/d
    
- a/+/c/d
    
- a/+/+/d
    
- +/+/+/+
    

The following subscriptions will not match:

- a/b/c
    
- b/+/c/d
    
- +/+/+
    

`#` can be used as a wildcard for all remaining levels of hierarchy. This means that it must be the final character in a subscription. With a topic of "a/b/c/d", the following example subscriptions will match:

- a/b/c/d
    
- #
    
- a/#
    
- a/b/#
    
- a/b/c/#
    
- +/b/c/#
    

Zero length topic levels are valid, which can lead to some slightly non-obvious behaviour. For example, a topic of "a//topic" would correctly match against a subscription of "a/+/topic". Likewise, zero length topic levels can exist at both the beginning and the end of a topic string, so "/a/topic" would match against a subscription of "+/a/topic", "#" or "/#", and a topic "a/topic/" would match against a subscription of "a/topic/+" or "a/topic/#".

### Quality of Service

MQTT defines three levels of Quality of Service (QoS). The QoS defines how hard the broker/client will try to ensure that a message is received. Messages may be sent at any QoS level, and clients may attempt to subscribe to topics at any QoS level. This means that the client chooses the maximum QoS it will receive. For example, if a message is published at QoS 2 and a client is subscribed with QoS 0, the message will be delivered to that client with QoS 0. If a second client is also subscribed to the same topic, but with QoS 2, then it will receive the same message but with QoS 2. For a second example, if a client is subscribed with QoS 2 and a message is published on QoS 0, the client will receive it on QoS 0.

Higher levels of QoS are more reliable, but involve higher latency and have higher bandwidth requirements.

- 0: The broker/client will deliver the message once, with no confirmation.
    
- 1: The broker/client will deliver the message at least once, with confirmation required.
    
- 2: The broker/client will deliver the message exactly once by using a four step handshake.
    

### Retained Messages

All messages may be set to be retained. This means that the broker will keep the message even after sending it to all current subscribers. If a new subscription is made that matches the topic of the retained message, then the message will be sent to the client. This is useful as a "last known good" mechanism. If a topic is only updated infrequently, then without a retained message, a newly subscribed client may have to wait a long time to receive an update. With a retained message, the client will receive an instant update.

### Clean session / Durable connections

On connection, a client sets the "clean session" flag, which is sometimes also known as the "clean start" flag. If clean session is set to false, then the connection is treated as durable. This means that when the client disconnects, any subscriptions it has will remain and any subsequent QoS 1 or 2 messages will be stored until it connects again in the future. If clean session is true, then all subscriptions will be removed for the client when it disconnects.

### Wills

When a client connects to a broker, it may inform the broker that it has a will. This is a message that it wishes the broker to send when the client disconnects unexpectedly. The will message has a topic, QoS and retain status just the same as any other message.