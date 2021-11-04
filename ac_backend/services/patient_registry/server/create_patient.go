package server

import (
	"alphacare/ac_backend/logging"
	pb "alphacare/ac_backend/proto/generated/patient_registry"
	user_registry_service "alphacare/ac_backend/proto/generated/user_registry"
	"alphacare/ac_backend/services/patient_registry/model"
	"alphacare/ac_backend/services/patient_registry/util"
	"context"
	"time"
)

var (
	log = logging.NewZapLogger()
)

func (s *Server) CreatePatient(ctx context.Context, in *pb.CreatPatientRequest) (*pb.CreatPatientResponse, error) {
	log.Info("Received: %v", in.GetPatient())
	res := &pb.CreatPatientResponse{}
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)

	patientId := util.GeneratePatientUUID()
	patient := &model.Patient{
		Id:                      patientId,
		FirstName:               in.GetPatient().GetFirstName(),
		LastName:                in.GetPatient().GetLastName(),
		SSN:                     in.GetPatient().GetSSN(),
		Gender:                  in.GetPatient().GetGender().String(),
		YearOfBirth:             in.GetPatient().GetYearOfBirth(),
		MonthOfBirth:            in.GetPatient().GetMonthOfBirth(),
		DayOfBirth:              in.GetPatient().GetDayOfBirth(),
		Status:                  in.GetPatient().GetStatus().String(),
		PrimaryCarePhysician:    in.GetPatient().GetPrimaryCarePhysician(),
		PhoneWork:               in.GetPatient().GetPhoneWork(),
		PhoneHome:               in.GetPatient().GetPhoneHome(),
		PhoneMobile:             in.GetPatient().GetPhoneMobile(),
		AddressStreet:           in.GetPatient().GetAddressStreet(),
		AddressCity:             in.GetPatient().GetAddressCity(),
		AddressState:            in.GetPatient().GetAddressState(),
		AddressZip:              in.GetPatient().GetAddressZip(),
		WeightInKg:              in.GetPatient().GetWeightInKg(),
		HeightInCm:              in.GetPatient().GetHeightInCm(),
		Ethnicity:               in.GetPatient().GetEthnicity(),
		Email:                   in.GetPatient().GetEmail(),
		PatientEmergencyContact: in.GetPatient().GetPatientEmergencyContact().String(),
		Referral:                in.GetPatient().GetReferral().String(),
		CreatedAt:               currentTime,
		UpdatedAt:               currentTime,
	}
	res.Patient = in.GetPatient()
	res.Patient.Id = patientId

	if err := s.db.Table(PatientTable).Create(patient).Error; err != nil {
		log.Errorf("Failed to create patient, error = %v", err)
		return nil, err
	}

	s.createUserEntryInUsersTable(ctx, patient)
	return res, nil
}

/*
this function calls createUser() endpoint to create an user entry in the users_table for this patient.
Password is autogenerated, the first time the patient logs into the application will need to reset password.
*/
func (s *Server) createUserEntryInUsersTable(ctx context.Context, patient *model.Patient) error {
	log.Infof("createUserEntryInUsersTable(): ")
	request := &user_registry_service.CreateUserRequest{
		FirstName:     patient.FirstName,
		LastName:      patient.LastName,
		Email:         patient.Email,
		Password:      util.GenerateInitialPassword(),
		ResetPassword: true, // TODO: notify the user with the default password. when log in, required to change password.
	}
	_, err := s.userRegistryClient.CreateUser(ctx, request)
	if err != nil {
		log.Errorf("failed to call CreateUser(), error = %v", err)
		return err
	}
	return nil
}