package main

import (
	pb "github.com/josemarjobs/shipper/consignment-service/proto/consignment"
	mgo "gopkg.in/mgo.v2"
)

const (
	dbName                = "shippy"
	consignmentCollection = "consignments"
)

type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

type ConsignmentRepo struct {
	session *mgo.Session
}

func (r *ConsignmentRepo) collection() *mgo.Collection {
	return r.session.DB(dbName).C(consignmentCollection)
}

func (r *ConsignmentRepo) Create(consignment *pb.Consignment) error {
	return r.collection().Insert(consignment)
}

func (r *ConsignmentRepo) GetAll() ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment
	err := r.collection().Find(nil).All(&consignments)
	return consignments, err
}

func (r *ConsignmentRepo) Close() {
	r.session.Close()
}
