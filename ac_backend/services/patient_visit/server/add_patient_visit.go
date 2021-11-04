package server

import (
	"alphacare/ac_backend/logging"
	pb "alphacare/ac_backend/proto/generated/patient_visit"
	patient_model "alphacare/ac_backend/services/patient_registry/model"
	"alphacare/ac_backend/services/patient_visit/model"
	"alphacare/ac_backend/services/patient_visit/util"
	user_model "alphacare/ac_backend/services/user_registry/model"
	"context"
	"strings"
)

var (
	log = logging.NewZapLogger()
)

func (s *Server) AddPatientVisit(ctx context.Context, in *pb.AddPatientVisitRequest) (*pb.AddPatientVisitResponse, error) {
	log.Info("AddPatientVisit for patient ID: %v", in.GetPatientId())
	res := &pb.AddPatientVisitResponse{}
	visitId := util.GeneratePatientVisitUUID()
	tmpUser := &user_model.User{}
	result := s.db.Table(UserTable).Where("id = ?", strings.ToLower(in.GetPatientId())).Find(tmpUser)
	if result.Error != nil {
		log.Info("User does not exist", in.GetPatientId())
		return res, result.Error
	}
	tmpPatient := &patient_model.Patient{}
	result = s.db.Table(PatientTable).Where("id = ?", strings.ToLower(in.GetPatientId())).Find(tmpPatient)
	if result.Error != nil {
		log.Info("Patient does not exist", in.GetPatientId())
		return res, result.Error
	}
	patientVisit := &model.PatientVisit{
		Id:        visitId,
		PatientId: in.GetPatientId(),
	}
	result = s.db.Table(PatientVisitTable).Create(patientVisit)
	if result.Error != nil {
		log.Info("Failed to add patient visit")
		return res, result.Error
	}

	res.Id = visitId
	return res, nil
}
