package internal

import (
	"errors"
	"strconv"
)

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) CreateUser(request CreateRequest) (*Membership, error) {

	for _, existingUser := range r.data {
		if existingUser.UserName == request.UserName {
			return nil, errors.New("이미 존재하는 이름입니다")
		}
	}

	if request.UserName == "" {
		return nil, errors.New("사용자 이름이 입력되지 않았습니다.")
	}

	if request.MembershipType == "" {
		return nil, errors.New("멤버십이 입력되지 않았습니다.")
	}

	whiteSlice := []string{"naver", "toss", "payco"}
	if !contains(whiteSlice, request.MembershipType) {
		return nil, errors.New("허용하지 않는 타입입니다.")
	}

	membership := Membership{
		ID:             strconv.Itoa(len(r.data) + 1),
		UserName:       request.UserName,
		MembershipType: request.MembershipType,
	}

	r.data[membership.ID] = membership
	return &membership, nil
}

func (r *Repository) UpdateUser(request UpdateRequest) (*Membership, error) {

	membership := r.data[request.ID]

	if membership.UserName != request.UserName {
		for _, existingUser := range r.data {
			if existingUser.UserName == request.UserName {
				return nil, errors.New("이미 존재하는 이름입니다")
			}
		}
	}

	if request.ID == "" {
		return nil, errors.New("멤버십 아이디 입력하지 않음")
	}

	if request.UserName == "" {
		return nil, errors.New("이름 입력하지 않음")
	}

	if request.MembershipType == "" {
		return nil, errors.New("멤버십을 입력하지 않음")
	}

	whiteSlice := []string{"naver", "toss", "payco"}
	if !contains(whiteSlice, request.MembershipType) {
		return nil, errors.New("허용하지 않는 타입니다")
	}

	newMembership := Membership{membership.ID, request.UserName, request.MembershipType}

	r.data[membership.ID] = newMembership
	return &newMembership, nil
}

func (r *Repository) DeleteUser(id string) error {

	if id == "" {
		return errors.New("id를 입력하지 않았습니다")
	}

	_, ok := r.data[id]
	if ok == false {
		return errors.New("입력한 id가 존재하지 않습니다")
	}

	delete(r.data, id)
	return nil
}

func contains(sliceValues []string, MembershipType string) bool {
	for _, value := range sliceValues {
		if value == MembershipType {
			return true
		}
	}
	return false
}
