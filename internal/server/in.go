package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindInHandler 查询所有 In 记录
func (s *Server) FindInHandler(c *gin.Context) {
	var Ins []models.In
	s.db.FindIns(&Ins)
	c.JSON(http.StatusOK, models.Message{In: Ins})
}

func (s *Server) DeleteInShouldIn(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldIn models.ShouldIn
	ShouldIn.ID = s.Str2Uint(c.PostForm("ShouldInID"))
	In.ShouldIns = append(In.ShouldIns, ShouldIn)

	if In.ID == 0 || ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 In 记录
	if err := s.db.DeleteInShouldIn(&In, &ShouldIn); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteInSend(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	var Send models.Send
	Send.ID = s.Str2Uint(c.PostForm("SendID"))
	In.Sends = append(In.Sends, Send)

	if In.ID == 0 || Send.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 In 记录
	if err := s.db.DeleteInSend(&In, &Send); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) FindInSendHandler(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	if In.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindInSend(&In)
	Send := In.Sends
	c.JSON(http.StatusOK, models.Message{Send: Send})
}

func (s *Server) FindInShouldInHandler(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	if In.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindInShouldIn(&In)
	ShouldIn := In.ShouldIns
	c.JSON(http.StatusOK, models.Message{ShouldIn: ShouldIn})
}

func (s *Server) FindInSaleHandler(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	s.db.FindInSale(&In)
	Sale := In.Sales
	c.JSON(http.StatusOK, models.Message{Sale: Sale})
}

// DeleteInHandler 删除 In 记录
func (s *Server) DeleteInHandler(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	if In.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 In: %+v\n", In)

	if err := s.db.Delete(&In); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteInSale(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	var Sale models.Sale
	Sale.ID = s.Str2Uint(c.PostForm("SaleID"))
	In.Sales = append(In.Sales, Sale)

	if In.ID == 0 || Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 In 记录
	if err := s.db.DeleteInSale(&In, &Sale); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) AddInSale(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	var Sale models.Sale
	Sale.ID = s.Str2Uint(c.PostForm("SaleID"))

	if In.ID == 0 || Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(In.ID, &In)
	s.db.FindByID(Sale.ID, &Sale)

	if Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	In.Sales = append(In.Sales, Sale)

	// 保存 In 记录
	if err := s.db.Save(In); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddInShouldIns(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldIn models.ShouldIn
	ShouldIn.ID = s.Str2Uint(c.PostForm("ShouldInID"))

	if In.ID == 0 || ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(In.ID, &In)
	s.db.FindByID(ShouldIn.ID, &ShouldIn)

	if ShouldIn.BillReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	In.ShouldIns = append(In.ShouldIns, ShouldIn)

	// 保存 In 记录
	if err := s.db.Save(In); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddInSends(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	var Send models.Send
	Send.ID = s.Str2Uint(c.PostForm("SendID"))

	if In.ID == 0 || Send.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(In.ID, &In)
	s.db.FindByID(Send.ID, &Send)

	if Send.SaleInvNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	In.Sends = append(In.Sends, Send)

	// 保存 In 记录
	if err := s.db.Save(In); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) SaveInHandler(c *gin.Context) {
	In := &models.In{}

	// 如果 ID 存在，查找已有记录
	In.ID = s.Str2Uint(c.PostForm("ID"))
	if In.ID != 0 {
		s.db.FindByID(In.ID, In)
	}
	In.MerchantID = s.Str2Uint(c.PostForm("MerchantID"))
	In.AcctID = s.Str2Uint(c.PostForm("AcctID"))
	In.BankAccountID = s.Str2Uint(c.PostForm("BankAccountID"))
	In.AcctBankID = s.Str2Uint(c.PostForm("AcctBankID"))

	// 验证必填字段
	if In.MerchantID == 0 || In.AcctID == 0 || In.BankAccountID == 0 || In.AcctBankID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败，必填字段缺失",
		})
		return
	}

	// 绑定嵌套结构体
	In.Merchant.ID = In.MerchantID
	In.Acct.ID = In.AcctID
	In.BankAccount.ID = In.BankAccountID
	In.AcctBank.ID = In.AcctBankID

	// 绑定其他字段
	In.ReceNum = c.PostForm("ReceNum")
	In.RealReceDate = c.PostForm("RealReceDate")
	In.ExpReceDate = c.PostForm("ExpReceDate")
	In.FinaDocType = c.PostForm("FinaDocType")
	In.FinaDocStatus = c.PostForm("FinaDocStatus")
	In.TotAmt = s.Str2Uint(c.PostForm("TotAmt"))
	In.Currency = c.PostForm("Currency")
	In.Notes = c.PostForm("Notes")
	In.FileID = s.Str2Uint(c.PostForm("FileID"))
	In.FileName = c.PostForm("FileName")

	log.Printf("保存 In: %+v\n", In)

	_, In.FileID, In.FileName = s.SaveFile(c, "file")
	// 保存 In 记录
	if err := s.db.SaveIn(In); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
