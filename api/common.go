package api

type resp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func success() resp {
	return resp{Message: "success"}
}

func successWithData(data interface{}) resp{
	return resp{Data: data}
}

func fail(msg string) resp {
	return resp{Message: msg}
}
