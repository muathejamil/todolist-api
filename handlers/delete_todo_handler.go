package handlers

// TodoDeleterSrv todo deleter service interface.
type TodoDeleterSrv interface {
	Delete(id uint) error
}

// DeleterSrv todo deleter service interface.
type DeleterSrv struct {
	serv TodoDeleterSrv
}

//// Delete deletes todo handler.
//// Params writer http.ResponseWriter, request *http.Request.
//// Returns.
//func (d *DeleterSrv) Delete(writer http.ResponseWriter, request *http.Request) {
//	id := utils.ExtractIdFromURL(writer, request)
//	err := d.serv.Delete(uint(id))
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusBadRequest)
//		return
//	}
//}
