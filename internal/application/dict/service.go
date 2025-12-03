package dict

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/domain/dict"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
)

type Service interface {
	// DictType
	CreateType(ctx context.Context, req *DictTypeRequest) error
	UpdateType(ctx context.Context, req *DictTypeRequest) error
	DeleteType(ctx context.Context, ids ...int64) error
	GetType(ctx context.Context, id int64) (*DictTypeResponse, error)
	ListType(ctx context.Context, req *DictTypeListRequest) (*DictTypeListResponse, error)

	// DictData
	CreateData(ctx context.Context, req *DictDataRequest) error
	UpdateData(ctx context.Context, req *DictDataRequest) error
	DeleteData(ctx context.Context, ids ...int64) error
	GetData(ctx context.Context, id int64) (*DictDataResponse, error)
	ListData(ctx context.Context, req *DictDataListRequest) (*DictDataListResponse, error)
}

type dictService struct {
	repo dict.Repository
}

func NewDictService(repo dict.Repository) Service {
	return &dictService{repo: repo}
}

// --- DictType ---

func (s *dictService) CreateType(ctx context.Context, req *DictTypeRequest) error {
	exist, _ := s.repo.FindTypeByCode(ctx, req.Code)
	if exist != nil {
		return errors.WithMsg(errors.AlreadyExists, "字典类型已存在")
	}
	return s.repo.CreateType(ctx, req.ToEntity())
}

func (s *dictService) UpdateType(ctx context.Context, req *DictTypeRequest) error {
	exist, err := s.repo.FindTypeByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if exist == nil {
		return errors.WithMsg(errors.NotFound, "字典类型不存在")
	}
	d := req.ToEntity()
	d.CreatedAt = exist.CreatedAt
	return s.repo.UpdateType(ctx, d)
}

func (s *dictService) DeleteType(ctx context.Context, ids ...int64) error {
	return s.repo.DeleteType(ctx, ids...)
}

func (s *dictService) GetType(ctx context.Context, id int64) (*DictTypeResponse, error) {
	d, err := s.repo.FindTypeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToDictTypeResponse(d), nil
}

func (s *dictService) ListType(ctx context.Context, req *DictTypeListRequest) (*DictTypeListResponse, error) {
	list, total, err := s.repo.ListType(ctx, req.ToQuery())
	if err != nil {
		return nil, err
	}
	var respList []*DictTypeResponse
	for _, d := range list {
		respList = append(respList, ToDictTypeResponse(d))
	}
	return &DictTypeListResponse{
		List:  respList,
		Total: total,
	}, nil
}

// --- DictData ---

func (s *dictService) CreateData(ctx context.Context, req *DictDataRequest) error {
	return s.repo.CreateData(ctx, req.ToEntity())
}

func (s *dictService) UpdateData(ctx context.Context, req *DictDataRequest) error {
	exist, err := s.repo.FindDataByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if exist == nil {
		return errors.WithMsg(errors.NotFound, "字典数据不存在")
	}
	d := req.ToEntity()
	d.CreatedAt = exist.CreatedAt
	return s.repo.UpdateData(ctx, d)
}

func (s *dictService) DeleteData(ctx context.Context, ids ...int64) error {
	return s.repo.DeleteData(ctx, ids...)
}

func (s *dictService) GetData(ctx context.Context, id int64) (*DictDataResponse, error) {
	d, err := s.repo.FindDataByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToDictDataResponse(d), nil
}

func (s *dictService) ListData(ctx context.Context, req *DictDataListRequest) (*DictDataListResponse, error) {
	list, total, err := s.repo.ListData(ctx, req.ToQuery())
	if err != nil {
		return nil, err
	}
	var respList []*DictDataResponse
	for _, d := range list {
		respList = append(respList, ToDictDataResponse(d))
	}
	return &DictDataListResponse{
		List:  respList,
		Total: total,
	}, nil
}
