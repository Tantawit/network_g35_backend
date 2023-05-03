package constant

// Chat Event

const ChatExchangeName = "chat"
const ChatExchangeKind = "topic"

// Queue Name convention <service>.<module>.event-type.<event type>.<event name>.queue

const ChatRegisterClientQueue = "chat.chat.event-type.topic.register.queue"
const ChatUnregisterClientQueue = "chat.chat.event-type.topic.unregister.queue"
const ChatBroadcastClientQueue = "chat.chat.event-type.topic.broadcast.queue"

// Topic Name convention <application>.<service>.<method>.<status>

const ChatRegisterClientTopicName = "cufreelance.chat.register-client.success"
const ChatUnregisterClientTopicName = "cufreelance.chat.unregister-client.success"
const ChatBroadcastTopicNameBase = "cufreelance.chat.broadcast"
