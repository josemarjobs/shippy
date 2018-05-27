package main

import (
	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	pb "github.com/josemarjobs/shipper/vessel-service/proto/vessel"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

type VesselRepo struct {
	session *mgo.Session
}

func (r *VesselRepo) collection() *mgo.Collection {
	return r.session.DB(dbName).C(vesselCollection)
}

func (repo *VesselRepo) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repo *VesselRepo) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

func (repo *VesselRepo) Close() {
	repo.session.Close()
}
