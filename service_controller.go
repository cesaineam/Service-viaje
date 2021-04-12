
package main

func createService(service Service) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO service (user_id,date_,state_service,vehicle_id,valor,cupos) VALUES (?, ?, ?,?,?,?)", service.UserId, service.Date, service.StateService,service.VehicleId,service.Valor,service.Cupos)
	return err
}

func deleteService(id int64) error {

	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM service WHERE service_id = ?", id)
	return err
}

// It takes the ID to make the update
func updateService(service Service) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE service SET user_id = ?, date_ = ?, state_service = ?, vehicle_id = ?, valor = ?, cupos = ? WHERE service_id = ?",service.UserId, service.Date, service.StateService,service.VehicleId,service.Valor,service.Cupos)
	return err
}
func getService() ([]Service, error) {
	//Declare an array because if there's error, we return it empty
	services := []Service{}
	bd, err := getDB()
	if err != nil {
		return services, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT service_id, user_id, date_, state_service, vehicle_id, valor, cupos FROM service")
	if err != nil {
		return services, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var service Service
		err = rows.Scan(&service.Id, &service.UserId, &service.Date, &service.StateService,&service.VehicleId,&service.Valor,&service.Cupos)
		if err != nil {
			return services, err
		}
		// and append it to the array
		services = append(services, service)
	}
	return services, nil
}

func getServiceById(id int64) (Service, error) {
	var service Service
	bd, err := getDB()
	if err != nil {
		return service, err
	}
	row := bd.QueryRow("SELECT service_id, user_id, date_, state_service, vehicle_id, valor, cupos FROM service WHERE service_id = ?", id)
	err = row.Scan(&service.Id, &service.UserId, &service.Date, &service.StateService,&service.VehicleId,&service.Valor,&service.Cupos)
	if err != nil {
		return service, err
	}
	// Success!
	return service, nil
}
