package controllers

import (
	"log"
	"net"

	pb "github.com/yerkebayev/go-final-go/proto"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func Controller(Port string, DB *gorm.DB) {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTeacherServiceServer(s, NewTeacherServer(DB))
	pb.RegisterStudentServiceServer(s, NewStudentServer(DB))
	pb.RegisterAdminServiceServer(s, NewAdminServer(DB))

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
