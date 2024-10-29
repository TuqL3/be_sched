package utils

type EquipmentStatus string

const (
	Working    EquipmentStatus = "working"
	Broken     EquipmentStatus = "broken"
	Maintained EquipmentStatus = "maintained"
)

type ReportStatus string

const (
	ReportPending  ReportStatus = "pending"
	ReportResolved ReportStatus = "resolved"
	ReportRejected ReportStatus = "rejected"
)

type RoomStatus string

const (
	RoomAvailable        RoomStatus = "available"
	RoomOccupied         RoomStatus = "occupied"
	RoomUnderMaintenance RoomStatus = "maintenance"
)

type ScheduleStatus string

const (
	ScheduleActive   ScheduleStatus = "active"
	ScheduleInactive ScheduleStatus = "inactive"
)
