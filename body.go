package engoBox2dSystem

import "github.com/ByteArena/box2d"

// Box2dComponent holds the box2d Body for use by Systems
type Box2dComponent struct {
	// Body is the box2d body
	Body *box2d.B2Body
}

var listOfBodiesToRemove []*box2d.B2Body

// DestroyBody destroys the box2d body from the World
// this does it safely at the end of an Update, so no bodies are removed during
// a simulation step, which can cause a crash
func (b *Box2dComponent) DestroyBody() {
	listOfBodiesToRemove = append(listOfBodiesToRemove, b.Body)
}

// removeBodies clears out all the box2d bodies on the list of bodies to remove.
// It is done separately, after DestroyBody is called so that no bodies are removed
// during a simulation step.
func removeBodies() {
	for _, bod := range listOfBodiesToRemove {
		World.DestroyBody(bod)
	}
	listOfBodiesToRemove = make([]*box2d.B2Body, 0)
}
