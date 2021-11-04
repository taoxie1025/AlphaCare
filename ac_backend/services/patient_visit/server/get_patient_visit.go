package server

import (
	pb "alphacare/ac_backend/proto/generated/patient_visit"
	"context"
	"strings"
)

func (s *Server) GetPatientVisits(ctx context.Context, req *pb.GetPatientVisitsRequest) (*pb.GetPatientVisitsResponse, error) {
	log.Info("GetPatientVisits: %v", req)
	res := &pb.GetPatientVisitsResponse{}
	//tmpVisits := &[]model.PatientVisit{}
	var tmpVisits []*pb.PatientVisit
	if req.GetVisitId() != "" {
		//result := s.db.Table(PatientVisitTable).Where("id = ?", strings.ToLower(req.GetVisitId())).Find(tmpVisits[0])
		result := s.db.Table(PatientVisitTable).Find(tmpVisits, req.GetVisitId())
		if result.Error != nil {
			return res, result.Error
		}
		return &pb.GetPatientVisitsResponse{
			PatientVisits: tmpVisits,
		}, nil
	}
	if req.GetPatientId() != "" {
		result := s.db.Table(PatientVisitTable).Where("patientId = ?", strings.ToLower(req.GetPatientId())).Find(tmpVisits)
		if result.Error != nil {
			return res, result.Error
		}
		return &pb.GetPatientVisitsResponse{
			PatientVisits: tmpVisits,
		}, nil
	}
	return res, nil
}
