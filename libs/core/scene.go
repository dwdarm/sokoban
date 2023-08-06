package core

type Scene interface {
	Start()
	Tick(input Input, timer Timer)
	Draw()
	Exit()
}

type SceneManager interface {
	SwitchScene(scene Scene)
	Start()
	Tick(input Input, timer Timer)
	Draw()
	Exit()
}

type SceneManagerImp struct {
	currentScene Scene
}

func NewSceneManager() SceneManager {
	return &SceneManagerImp{}
}

func (sm *SceneManagerImp) SwitchScene(scene Scene) {
	if sm.currentScene != nil {
		sm.currentScene.Exit()
	}

	sm.currentScene = scene
	sm.currentScene.Start()
}

func (sm *SceneManagerImp) Start() {
	if sm.currentScene != nil {
		sm.currentScene.Start()
	}
}

func (sm *SceneManagerImp) Tick(input Input, timer Timer) {
	if sm.currentScene != nil {
		sm.currentScene.Tick(input, timer)
	}
}

func (sm *SceneManagerImp) Draw() {
	if sm.currentScene != nil {
		sm.currentScene.Draw()
	}
}

func (sm *SceneManagerImp) Exit() {
	if sm.currentScene != nil {
		sm.currentScene.Exit()
	}
}
