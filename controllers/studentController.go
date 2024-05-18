package controllers

import (
	"context"
	"github.com/yerkebayev/go-final-go/models"
	pb "github.com/yerkebayev/go-final-go/proto"
	"gorm.io/gorm"
	"time"
)

type StudentServer struct {
	pb.UnimplementedStudentServiceServer
	db *gorm.DB
}

func NewStudentServer(db *gorm.DB) *StudentServer {
	return &StudentServer{db: db}
}

func (s *StudentServer) AddStudent(ctx context.Context, req *pb.AddStudentRequest) (*pb.AddStudentResponse, error) {
	student := models.Student{
		StudentNumberId: req.StudentNumberId,
		Name:            req.Name,
	}

	if err := s.db.Create(&student).Error; err != nil {
		return nil, err
	}

	return &pb.AddStudentResponse{Id: int32(student.ID)}, nil
}

func (s *StudentServer) AddAttendance(ctx context.Context, req *pb.AttendanceRequest) (*pb.AttendanceResponse, error) {
	attendance := models.Attendance{
		SessionID: uint(req.SessionId),
		StudentID: uint(req.StudentId),
		Time:      time.Now().Format(time.RFC3339),
	}
	if err := s.db.Create(&attendance).Error; err != nil {
		return nil, err
	}
	return &pb.AttendanceResponse{Id: int32(attendance.ID)}, nil
}

func (s *StudentServer) UploadImages(ctx context.Context, req *pb.UploadImagesRequest) (*pb.UploadImagesResponse, error) {
	for _, img := range req.Images {
		image := models.Image{
			UserID:    int(req.UserId),
			ImageData: img.Data,
		}
		s.db.Create(&image)
	}
	return &pb.UploadImagesResponse{Status: "success"}, nil
}
func (s *StudentServer) GetImages(ctx context.Context, req *pb.GetImagesRequest) (*pb.GetImagesResponse, error) {
	var images []models.Image
	s.db.Find(&images)

	var pbImages []*pb.Image
	for _, img := range images {
		pbImages = append(pbImages, &pb.Image{
			Id:     int32(img.ID),
			UserId: int32(img.UserID),
			Data:   img.ImageData,
		})
	}

	return &pb.GetImagesResponse{Images: pbImages}, nil
}

func (s *StudentServer) GetStudentNames(ctx context.Context, req *pb.GetStudentNamesRequest) (*pb.GetStudentNamesResponse, error) {
	var students []models.Student
	s.db.Find(&students)

	var pbStudents []*pb.Student
	for _, student := range students {
		pbStudents = append(pbStudents, &pb.Student{
			Id:              int32(student.ID),
			StudentNumberId: student.StudentNumberId,
			Name:            student.Name,
		})
	}

	return &pb.GetStudentNamesResponse{Students: pbStudents}, nil
}
