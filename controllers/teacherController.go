package controllers

import (
	"context"

	"github.com/yerkebayev/go-final-go/models"
	pb "github.com/yerkebayev/go-final-go/proto"
	"gorm.io/gorm"
)

type TeacherServer struct {
	pb.UnimplementedTeacherServiceServer
	db *gorm.DB
}

func NewTeacherServer(db *gorm.DB) *TeacherServer {
	return &TeacherServer{db: db}
}

func (s *TeacherServer) RegisterTeacher(ctx context.Context, req *pb.RegisterTeacherRequest) (*pb.RegisterTeacherResponse, error) {
	teacher := models.Teacher{Name: req.Name}
	if err := s.db.Create(&teacher).Error; err != nil {
		return nil, err
	}
	return &pb.RegisterTeacherResponse{Id: int32(teacher.ID)}, nil
}

func (s *TeacherServer) AddSession(ctx context.Context, req *pb.AddSessionRequest) (*pb.AddSessionResponse, error) {
	session := models.Session{
		TeacherID: uint(req.TeacherId),
		CourseID:  uint(req.CourseId),
		Date:      req.Date,
	}
	if err := s.db.Create(&session).Error; err != nil {
		return nil, err
	}
	return &pb.AddSessionResponse{Id: int32(session.ID)}, nil
}

func (s *TeacherServer) GetSession(ctx context.Context, req *pb.GetSessionRequest) (*pb.GetSessionResponse, error) {
	var session models.Session
	if err := s.db.Preload("Attendances").First(&session, req.Id).Error; err != nil {
		return nil, err
	}

	attendances := make([]*pb.Attendance, len(session.Attendances))
	for i, attendance := range session.Attendances {
		attendances[i] = &pb.Attendance{
			Id:        int32(attendance.ID),
			SessionId: int32(attendance.SessionID),
			StudentId: int32(attendance.StudentID),
			Time:      attendance.Time,
		}
	}

	return &pb.GetSessionResponse{
		Id:          int32(session.ID),
		TeacherId:   int32(session.TeacherID),
		CourseId:    int32(session.CourseID),
		Date:        session.Date,
		Attendances: attendances,
	}, nil
}
