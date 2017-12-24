package httpapi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/talento90/gorpo/pkg/gorpo"
)

type effectsController struct {
	service gorpo.EffectService
}

func newEffectsController(service gorpo.EffectService) *effectsController {
	return &effectsController{
		service: service,
	}
}

func (c *effectsController) GetEffectByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) appResponse {
	id := params.ByName("id")

	effect, err := c.service.GetEffect(id)

	if err != nil {
		return errResponse(err)
	}

	return response(http.StatusOK, effect)
}

func (c *effectsController) GetAllEffects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) appResponse {
	effects, err := c.service.GetEffects()

	if err != nil {
		return errResponse(err)
	}

	return response(http.StatusOK, effects)
}
