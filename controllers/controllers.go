package controllers

import (
	"context"
	"github.com/yerkebayev/go-final-go/models"
	"log"
	"net"

	pb "github.com/yerkebayev/go-final-go/proto"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type TeacherServer struct {
	pb.UnimplementedTeacherServiceServer
	db *gorm.DB
}

func NewTeacherServer(db *gorm.DB) *TeacherServer {
	return &TeacherServer{db: db}
}

func (s *TeacherServer) GetReport(ctx context.Context, req *pb.TeacherRequest) (*pb.TeacherReportResponse, error) {
	// Implement your logic here
	return &pb.TeacherReportResponse{Reports: []string{"Report 1", "Report 2"}}, nil
}

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

func (s *StudentServer) AddStudentToCourse(ctx context.Context, req *pb.AddStudentToCourseRequest) (*pb.AddStudentToCourseResponse, error) {
	// Implement your logic here
	return &pb.AddStudentToCourseResponse{Success: true, Message: "Student added to course"}, nil
}

type MainServer struct {
	pb.UnimplementedMainServiceServer
	db *gorm.DB
}

func NewMainServer(db *gorm.DB) *MainServer {
	return &MainServer{db: db}
}

func (s *MainServer) AddAttendance(ctx context.Context, req *pb.AttendanceRequest) (*pb.AttendanceResponse, error) {
	// Implement your logic here
	return &pb.AttendanceResponse{Id: req.Id, Details: "Attendance details"}, nil
}

type ImageServer struct {
	pb.UnimplementedImageServiceServer
	DB *gorm.DB
}

func NewImageServer(db *gorm.DB) *ImageServer {
	return &ImageServer{DB: db}
}

func (s *ImageServer) UploadImages(ctx context.Context, req *pb.UploadImagesRequest) (*pb.UploadImagesResponse, error) {
	for _, img := range req.Images {
		image := models.Image{
			UserID:    int(req.UserId),
			ImageData: img.Data,
		}
		s.DB.Create(&image)
	}
	return &pb.UploadImagesResponse{Status: "success"}, nil
}
func (s *ImageServer) GetImages(ctx context.Context, req *pb.GetImagesRequest) (*pb.GetImagesResponse, error) {
	var images []models.Image
	s.DB.Find(&images)

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

func Controller(Port string, DB *gorm.DB) {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTeacherServiceServer(s, NewTeacherServer(DB))
	pb.RegisterStudentServiceServer(s, NewStudentServer(DB))
	pb.RegisterMainServiceServer(s, NewMainServer(DB))
	pb.RegisterImageServiceServer(s, NewImageServer(DB))

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
