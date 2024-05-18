package controllers

import (
	"context"
	"github.com/yerkebayev/go-final-go/models"
	pb "github.com/yerkebayev/go-final-go/proto"
	"gorm.io/gorm"
)

type AdminServer struct {
	pb.UnimplementedAdminServiceServer
	DB *gorm.DB
}

func NewAdminServer(db *gorm.DB) *AdminServer {
	return &AdminServer{DB: db}
}

func (s *AdminServer) AddStudent(ctx context.Context, req *pb.AddStudentRequest) (*pb.AddStudentResponse, error) {
	student := models.Student{
		StudentNumberId: req.StudentNumberId,
		Name:            req.Name,
	}
	s.DB.Create(&student)
	return &pb.AddStudentResponse{Id: int32(student.ID)}, nil
}

func (s *AdminServer) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentResponse, error) {
	if err := s.DB.Delete(&models.Student{}, req.Id).Error; err != nil {
		return nil, err
	}
	return &pb.DeleteStudentResponse{Id: req.Id}, nil
}

func (s *AdminServer) DeleteTeacher(ctx context.Context, req *pb.DeleteTeacherRequest) (*pb.DeleteTeacherResponse, error) {
	if err := s.DB.Delete(&models.Teacher{}, req.Id).Error; err != nil {
		return nil, err
	}
	return &pb.DeleteTeacherResponse{Id: req.Id}, nil
}

func (s *AdminServer) AddCourse(ctx context.Context, req *pb.AddCourseRequest) (*pb.AddCourseResponse, error) {
	course := models.Course{
		Title: req.Name,
	}
	s.DB.Create(&course)
	return &pb.AddCourseResponse{Id: int32(course.ID)}, nil
}

func (s *AdminServer) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	if err := s.DB.Delete(&models.Course{}, req.Id).Error; err != nil {
		return nil, err
	}
	return &pb.DeleteCourseResponse{Id: req.Id}, nil
}
func (s *AdminServer) GetCourses(ctx context.Context, req *pb.Empty) (*pb.GetCoursesResponse, error) {
	var courses []models.Course
	if err := s.DB.Find(&courses).Error; err != nil {
		return nil, err
	}

	var pbCourses []*pb.Course
	for _, course := range courses {
		pbCourses = append(pbCourses, &pb.Course{
			Id:    int32(course.ID),
			Title: course.Title,
		})
	}
	return &pb.GetCoursesResponse{Courses: pbCourses}, nil
}

func (s *AdminServer) GetStudents(ctx context.Context, req *pb.Empty) (*pb.GetStudentsResponse, error) {
	var students []models.Student
	if err := s.DB.Find(&students).Error; err != nil {
		return nil, err
	}

	var pbStudents []*pb.Student
	for _, student := range students {
		pbStudents = append(pbStudents, &pb.Student{
			Id:              int32(student.ID),
			StudentNumberId: student.StudentNumberId,
			Name:            student.Name,
		})
	}
	return &pb.GetStudentsResponse{Students: pbStudents}, nil
}

func (s *AdminServer) GetTeachers(ctx context.Context, req *pb.Empty) (*pb.GetTeachersResponse, error) {
	var teachers []models.Teacher
	if err := s.DB.Find(&teachers).Error; err != nil {
		return nil, err
	}

	var pbTeachers []*pb.Teacher
	for _, teacher := range teachers {
		pbTeachers = append(pbTeachers, &pb.Teacher{
			Id:   int32(teacher.ID),
			Name: teacher.Name,
		})
	}
	return &pb.GetTeachersResponse{Teachers: pbTeachers}, nil
}

//func (s *AdminServer) AddStudentToCourse(ctx context.Context, req *pb.AddStudentToCourseRequest) (*pb.AddStudentToCourseResponse, error) {
//	studentCourse := models.StudentCourse{
//		StudentID: int(req.StudentId),
//		CourseID:  int(req.CourseId),
//	}
//	if err := s.DB.Create(&studentCourse).Error; err != nil {
//		return nil, err
//	}
//	return &pb.AddStudentToCourseResponse{Success: true}, nil
//}
