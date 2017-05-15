// THIS IS AN AUTOMATICALLY GENERATED CODE! DO NOT EDIT THIS FILE!
// ADD YOUR EVENT TO 'generate_event_dispatcher.go' AND RUN 'go generate'

package events

import (
	"time"
)

// #######################
// INTERFACE DOCUMENTATION
// #######################

// 1. Create EventDispatcher using NewEventDispatcher() function.
// 2. Register your listeners using EventDispatcher.Register<event type name>Listener methods.
// 3. Run event loop by calling EventDispatcher.RunEventLoop() method.
// 4. Trigger events using EventDispatcher.Fire<event type name> methods.

// LISTENER INTERFACES

type TimeTickListener interface {
	HandleTimeTick(*TimeTick)
}

type PhysicsReadyListener interface {
	HandlePhysicsReady(*PhysicsReady)
}

type ProjectileFiredListener interface {
	HandleProjectileFired(*ProjectileFired)
}

type ProjectileHitListener interface {
	HandleProjectileHit(*ProjectileHit)
}

type UserConnectedListener interface {
	HandleUserConnected(*UserConnected)
}

type UserJoinedListener interface {
	HandleUserJoined(*UserJoined)
}

type UserLeftListener interface {
	HandleUserLeft(*UserLeft)
}

type ObjectDestroyedListener interface {
	HandleObjectDestroyed(*ObjectDestroyed)
}

type UserInputListener interface {
	HandleUserInput(*UserInput)
}

type TargetAngleListener interface {
	HandleTargetAngle(*TargetAngle)
}

type ScoreSentListener interface {
	HandleScoreSent(*ScoreSent)
}

// ##############################
// END OF INTERFACE DOCUMENTATION
// ##############################

const (
	eventQueuesCapacity                                       = 100000
	idleDispatcherSleepTime                     time.Duration = 5 * time.Millisecond
	registeringListenerWhileRunningErrorMessage               = "Tried to register listener while running event loop. Registering listeners is not thread safe therefore prohibited after starting event loop."
)

// PRIVATE EVENT HANDLERS

type eventHandler interface {
	handle()
}

type timeTickHandler struct {
	event          *TimeTick
	eventListeners []TimeTickListener
}

func (handler *timeTickHandler) handle() {
	for _, listener := range handler.eventListeners {
		listener.HandleTimeTick(handler.event)
	}
}

type physicsReadyHandler struct {
	event          *PhysicsReady
	eventListeners []PhysicsReadyListener
}

func (handler *physicsReadyHandler) handle() {
	for _, listener := range handler.eventListeners {
		listener.HandlePhysicsReady(handler.event)
	}
}

type projectileFiredHandler struct {
	event          *ProjectileFired
	eventListeners []ProjectileFiredListener
}

func (handler *projectileFiredHandler) handle() {
	for _, listener := range handler.eventListeners {
		listener.HandleProjectileFired(handler.event)
	}
}

type projectileHitHandler struct {
	event          *ProjectileHit
	eventListeners []ProjectileHitListener
}

func (handler *projectileHitHandler) handle() {
	for _, listener := range handler.eventListeners {
		listener.HandleProjectileHit(handler.event)
	}
}

type userConnectedHandler struct {
	event          *UserConnected
	eventListeners []UserConnectedListener
}

func (handler *userConnectedHandler) handle() {
	for _, listener := range handler.eventListeners {
		listener.HandleUserConnected(handler.event)
	}
}

type userJoinedHandler struct {
	event          *UserJoined
	eventListeners []UserJoinedListener
}

func (handler *userJoinedHandler) handle() {
	for _, listener := range handler.eventListeners {
		listener.HandleUserJoined(handler.event)
	}
}

type userLeftHandler struct {
	event          *UserLeft
	eventListeners []UserLeftListener
}

func (handler *userLeftHandler) handle() {
	for _, listener := range handler.eventListeners {
		listener.HandleUserLeft(handler.event)
	}
}

type objectDestroyedHandler struct {
	event          *ObjectDestroyed
	eventListeners []ObjectDestroyedListener
}

func (handler *objectDestroyedHandler) handle() {
	for _, listener := range handler.eventListeners {
		listener.HandleObjectDestroyed(handler.event)
	}
}

type userInputHandler struct {
	event          *UserInput
	eventListeners []UserInputListener
}

func (handler *userInputHandler) handle() {
	for _, listener := range handler.eventListeners {
		listener.HandleUserInput(handler.event)
	}
}

type targetAngleHandler struct {
	event          *TargetAngle
	eventListeners []TargetAngleListener
}

func (handler *targetAngleHandler) handle() {
	for _, listener := range handler.eventListeners {
		listener.HandleTargetAngle(handler.event)
	}
}

type scoreSentHandler struct {
	event          *ScoreSent
	eventListeners []ScoreSentListener
}

func (handler *scoreSentHandler) handle() {
	for _, listener := range handler.eventListeners {
		listener.HandleScoreSent(handler.event)
	}
}

// EVENT DISPATCHER

type EventDispatcher struct {
	running bool

	// EVENT QUEUES

	priority1EventsQueue chan eventHandler

	priority2EventsQueue chan eventHandler

	priority3EventsQueue chan eventHandler

	// LISTENER LISTS

	timeTickListeners []TimeTickListener

	physicsReadyListeners []PhysicsReadyListener

	projectileFiredListeners []ProjectileFiredListener

	projectileHitListeners []ProjectileHitListener

	userConnectedListeners []UserConnectedListener

	userJoinedListeners []UserJoinedListener

	userLeftListeners []UserLeftListener

	objectDestroyedListeners []ObjectDestroyedListener

	userInputListeners []UserInputListener

	targetAngleListeners []TargetAngleListener

	scoreSentListeners []ScoreSentListener
}

// EVENT DISPATCHER CONSTRUCTOR

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		running: false,

		// EVENT QUEUES

		priority1EventsQueue: make(chan eventHandler, eventQueuesCapacity),

		priority2EventsQueue: make(chan eventHandler, eventQueuesCapacity),

		priority3EventsQueue: make(chan eventHandler, eventQueuesCapacity),

		// LISTENER LISTS

		timeTickListeners: []TimeTickListener{},

		physicsReadyListeners: []PhysicsReadyListener{},

		projectileFiredListeners: []ProjectileFiredListener{},

		projectileHitListeners: []ProjectileHitListener{},

		userConnectedListeners: []UserConnectedListener{},

		userJoinedListeners: []UserJoinedListener{},

		userLeftListeners: []UserLeftListener{},

		objectDestroyedListeners: []ObjectDestroyedListener{},

		userInputListeners: []UserInputListener{},

		targetAngleListeners: []TargetAngleListener{},

		scoreSentListeners: []ScoreSentListener{},
	}
}

// MAIN EVENT LOOP

func (dispatcher *EventDispatcher) RunEventLoop() {
	dispatcher.running = true

	for {
		select {

		case handler := <-dispatcher.priority1EventsQueue:
			handler.handle()

		case handler := <-dispatcher.priority2EventsQueue:
			handler.handle()

		case handler := <-dispatcher.priority3EventsQueue:
			handler.handle()

		default:
			time.Sleep(idleDispatcherSleepTime)
		}
	}
}

func (dispatcher *EventDispatcher) panicWhenEventLoopRunning() {
	if dispatcher.running {
		panic(registeringListenerWhileRunningErrorMessage)
	}
}

// PUBLIC EVENT DISPATCHER METHODS

type QueueFilling struct {
	CurrentLength int
	Capacity      int
}

func (dispatcher *EventDispatcher) QueuesFilling() map[int]QueueFilling {
	filling := make(map[int]QueueFilling)

	filling[1] = QueueFilling{len(dispatcher.priority1EventsQueue), eventQueuesCapacity}

	filling[2] = QueueFilling{len(dispatcher.priority2EventsQueue), eventQueuesCapacity}

	filling[3] = QueueFilling{len(dispatcher.priority3EventsQueue), eventQueuesCapacity}

	return filling
}

// TimeTick

func (dispatcher *EventDispatcher) RegisterTimeTickListener(listener TimeTickListener) {
	dispatcher.panicWhenEventLoopRunning()

	dispatcher.timeTickListeners = append(dispatcher.timeTickListeners, listener)
}

func (dispatcher *EventDispatcher) FireTimeTick(event *TimeTick) {
	handler := &timeTickHandler{
		event:          event,
		eventListeners: dispatcher.timeTickListeners,
	}

	dispatcher.priority1EventsQueue <- handler
}

// PhysicsReady

func (dispatcher *EventDispatcher) RegisterPhysicsReadyListener(listener PhysicsReadyListener) {
	dispatcher.panicWhenEventLoopRunning()

	dispatcher.physicsReadyListeners = append(dispatcher.physicsReadyListeners, listener)
}

func (dispatcher *EventDispatcher) FirePhysicsReady(event *PhysicsReady) {
	handler := &physicsReadyHandler{
		event:          event,
		eventListeners: dispatcher.physicsReadyListeners,
	}

	dispatcher.priority1EventsQueue <- handler
}

// ProjectileFired

func (dispatcher *EventDispatcher) RegisterProjectileFiredListener(listener ProjectileFiredListener) {
	dispatcher.panicWhenEventLoopRunning()

	dispatcher.projectileFiredListeners = append(dispatcher.projectileFiredListeners, listener)
}

func (dispatcher *EventDispatcher) FireProjectileFired(event *ProjectileFired) {
	handler := &projectileFiredHandler{
		event:          event,
		eventListeners: dispatcher.projectileFiredListeners,
	}

	dispatcher.priority2EventsQueue <- handler
}

// ProjectileHit

func (dispatcher *EventDispatcher) RegisterProjectileHitListener(listener ProjectileHitListener) {
	dispatcher.panicWhenEventLoopRunning()

	dispatcher.projectileHitListeners = append(dispatcher.projectileHitListeners, listener)
}

func (dispatcher *EventDispatcher) FireProjectileHit(event *ProjectileHit) {
	handler := &projectileHitHandler{
		event:          event,
		eventListeners: dispatcher.projectileHitListeners,
	}

	dispatcher.priority2EventsQueue <- handler
}

// UserConnected

func (dispatcher *EventDispatcher) RegisterUserConnectedListener(listener UserConnectedListener) {
	dispatcher.panicWhenEventLoopRunning()

	dispatcher.userConnectedListeners = append(dispatcher.userConnectedListeners, listener)
}

func (dispatcher *EventDispatcher) FireUserConnected(event *UserConnected) {
	handler := &userConnectedHandler{
		event:          event,
		eventListeners: dispatcher.userConnectedListeners,
	}

	dispatcher.priority2EventsQueue <- handler
}

// UserJoined

func (dispatcher *EventDispatcher) RegisterUserJoinedListener(listener UserJoinedListener) {
	dispatcher.panicWhenEventLoopRunning()

	dispatcher.userJoinedListeners = append(dispatcher.userJoinedListeners, listener)
}

func (dispatcher *EventDispatcher) FireUserJoined(event *UserJoined) {
	handler := &userJoinedHandler{
		event:          event,
		eventListeners: dispatcher.userJoinedListeners,
	}

	dispatcher.priority2EventsQueue <- handler
}

// UserLeft

func (dispatcher *EventDispatcher) RegisterUserLeftListener(listener UserLeftListener) {
	dispatcher.panicWhenEventLoopRunning()

	dispatcher.userLeftListeners = append(dispatcher.userLeftListeners, listener)
}

func (dispatcher *EventDispatcher) FireUserLeft(event *UserLeft) {
	handler := &userLeftHandler{
		event:          event,
		eventListeners: dispatcher.userLeftListeners,
	}

	dispatcher.priority2EventsQueue <- handler
}

// ObjectDestroyed

func (dispatcher *EventDispatcher) RegisterObjectDestroyedListener(listener ObjectDestroyedListener) {
	dispatcher.panicWhenEventLoopRunning()

	dispatcher.objectDestroyedListeners = append(dispatcher.objectDestroyedListeners, listener)
}

func (dispatcher *EventDispatcher) FireObjectDestroyed(event *ObjectDestroyed) {
	handler := &objectDestroyedHandler{
		event:          event,
		eventListeners: dispatcher.objectDestroyedListeners,
	}

	dispatcher.priority2EventsQueue <- handler
}

// UserInput

func (dispatcher *EventDispatcher) RegisterUserInputListener(listener UserInputListener) {
	dispatcher.panicWhenEventLoopRunning()

	dispatcher.userInputListeners = append(dispatcher.userInputListeners, listener)
}

func (dispatcher *EventDispatcher) FireUserInput(event *UserInput) {
	handler := &userInputHandler{
		event:          event,
		eventListeners: dispatcher.userInputListeners,
	}

	dispatcher.priority3EventsQueue <- handler
}

// TargetAngle

func (dispatcher *EventDispatcher) RegisterTargetAngleListener(listener TargetAngleListener) {
	dispatcher.panicWhenEventLoopRunning()

	dispatcher.targetAngleListeners = append(dispatcher.targetAngleListeners, listener)
}

func (dispatcher *EventDispatcher) FireTargetAngle(event *TargetAngle) {
	handler := &targetAngleHandler{
		event:          event,
		eventListeners: dispatcher.targetAngleListeners,
	}

	dispatcher.priority3EventsQueue <- handler
}

// ScoreSent

func (dispatcher *EventDispatcher) RegisterScoreSentListener(listener ScoreSentListener) {
	dispatcher.panicWhenEventLoopRunning()

	dispatcher.scoreSentListeners = append(dispatcher.scoreSentListeners, listener)
}

func (dispatcher *EventDispatcher) FireScoreSent(event *ScoreSent) {
	handler := &scoreSentHandler{
		event:          event,
		eventListeners: dispatcher.scoreSentListeners,
	}

	dispatcher.priority3EventsQueue <- handler
}
