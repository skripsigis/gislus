package commons

func GetMenus() []string {
	menus := []string{}

	menus = append(menus, "dashboard")
	menus = append(menus, "master")
	menus = append(menus, "master_employee")
	menus = append(menus, "master_designation")
	menus = append(menus, "master_department")
	menus = append(menus, "master_location")
	menus = append(menus, "master_competency")
	menus = append(menus, "master_measurement")
	menus = append(menus, "master_role")
	menus = append(menus, "master_dateprocess")
	menus = append(menus, "master_menu")
	menus = append(menus, "goalsetting_menu")
	menus = append(menus, "assessment")
	menus = append(menus, "assessment_approval")
	menus = append(menus, "competency_assessment")
	menus = append(menus, "fact_gathering")
	menus = append(menus, "fact_gathering_personal")
	menus = append(menus, "fact_gathering_subordinate")
	menus = append(menus, "notif")
	menus = append(menus, "notif_info")
	menus = append(menus, "notif_history")
	menus = append(menus, "history")
	menus = append(menus, "assessment_history")
	menus = append(menus, "history_competency_assessment")
	menus = append(menus, "notif_goal_setting")
	menus = append(menus, "notif_pending")
	menus = append(menus, "notif_pending_assessor")
	menus = append(menus, "review")
	menus = append(menus, "midyearreview")
	menus = append(menus, "hodreview")
	menus = append(menus, "endyearreview")
	menus = append(menus, "endyearreview_emp")
	menus = append(menus, "endyearreview_hod")
	menus = append(menus, "competency_assessment")
	menus = append(menus, "competency_assessment_emp")
	menus = append(menus, "competency_assessment_hod")

	return menus
}

// func GetMenusByUser(roleid string) map[string]bool {
// 	//menus := GetMenus()
// 	retMenus := make(map[string]bool)
// 	// found := false
// 	// for k, v := range menus {
// 	// 	for _, rm := range v {
// 	// 		for _, rp := range roles {
// 	// 			if rp == rm {
// 	// 				found = true
// 	// 			}
// 	// 		}
// 	// 	}
// 	// 	retMenus[k] = found

// 	// 	found = false
// 	// }

// 	return retMenus
// }

// func GetAccessMenu(roleid string, menu string) bool {
// 	if menu == "dashboard" {
// 		return true
// 	}

// 	menus := GetMenus()
// 	menuRoles := menus[menu]

// 	for _, rp := range roles {
// 		for _, rm := range menuRoles {
// 			if rm == rp {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }
