package server

import (
	"alphacare/ac_backend/logging"
	pb "alphacare/ac_backend/proto/generated/prescription"
	"alphacare/ac_backend/services/prescription/util"
	"context"
	"time"
)

var (
	log = logging.NewZapLogger()
)

func (s *Server) AddPrescription(ctx context.Context, in *pb.AddPrescriptionRequest) (*pb.AddPrescriptionResponse, error) {
	log.Info("AddPrescription for patient Id: %v ordered by: ", in.GetPrescription().GetPatientId(), in.GetPrescription().GetOrderedBy())

	res := &pb.AddPrescriptionResponse{}
	patientId := in.GetPrescription().GetPatientId()
	orderedBy := in.GetPrescription().GetOrderedBy()
	if patientId == "" {
		log.Info("Patient can not be null")
		return res, nil
	}
	if orderedBy == "" {
		log.Info("Ordered by can not be null")
		return res, nil
	}
	prescription := in.GetPrescription()
	prescription.Id = util.GeneratePrescriptionUUID()
	prescription.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	result := s.db.Table(PrescriptionTable).Create(prescription)
	if result.Error != nil {
		log.Info("Failed to add prescription")
		return res, result.Error
	}
	return res, nil
}
