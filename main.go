package main

import (
	"context"
	"log"
	"net"

	pb "github.com/hlb8122/uki-track-backend/management"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

/*
Methods for exercise management.
*/
type managementServer struct{}

func (s *managementServer) AddLiftType(ctx context.Context, in *pb.String) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *managementServer) RemoveLiftType(ctx context.Context, in *pb.String) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *managementServer) ListLiftTypes(ctx context.Context, in *pb.Empty) (*pb.ListString, error) {
	listLifts := []string{"bench", "lat pulldown", "squat"}
	return &pb.ListString{Value: listLifts}, nil
}

func (s *managementServer) AddCardioType(ctx context.Context, in *pb.String) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *managementServer) RemoveCardioType(ctx context.Context, in *pb.String) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *managementServer) ListCardioType(ctx context.Context, in *pb.Empty) (*pb.ListString, error) {
	listCardio := []string{"bench", "lat pulldown", "squat"}
	return &pb.ListString{Value: listCardio}, nil
}

/*
Methods for routine management.
*/

func (s *managementServer) AddEmptyRoutine(ctx context.Context, in *pb.String) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *managementServer) CloneRoutine(ctx context.Context, in *pb.String) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *managementServer) RemoveRoutine(ctx context.Context, in *pb.String) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *managementServer) ListRoutines(ctx context.Context, in *pb.Empty) (*pb.ListString, error) {
	listCardio := []string{"bench", "lat pulldown", "squat"}
	return &pb.ListString{Value: listCardio}, nil
}

func (s *managementServer) InsertWorkout(ctx context.Context, in *pb.RoutineWorkout) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

/*
Methods for workout management.
*/

func (s *managementServer) AddEmptyWorkout(ctx context.Context, in *pb.String) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *managementServer) RemoveWorkout(ctx context.Context, in *pb.String) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *managementServer) ListWorkouts(ctx context.Context, in *pb.Empty) (*pb.ListString, error) {
	listCardio := []string{"bench", "lat pulldown", "squat"}
	return &pb.ListString{Value: listCardio}, nil
}

func (s *managementServer) InsertSet(ctx context.Context, in *pb.WorkoutSet) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterExerciseManagementServer(s, &managementServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
