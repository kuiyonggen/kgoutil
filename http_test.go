package kgoutil

import (
	"testing"
)


type Symbol struct {
    AStockCode      string `json:"a_stock_code"     example:"600010"`
    CompanyAbbr     string `json:"company_abbr"     example:"123456"`
    FullName        string `json:"full_name"        example:"full name"`
    ListDate        string `json:"list_date"        example:"list date"`
    CSRCCode        string `json:"csrc_code"        example:"F"`
    CSRCCodeDesc    string `json:"csrc_code_desc"   example:"csrc code desc"`
}

type responseData struct {
    Result []Symbol `json:"result"`
}

func Test_Get(t *testing.T) {
	tests := []struct {
		Name string
                Url  string
                Headers map[string]string
                Params map[string]string
	}{
		{
                    "Test_Get",
                    "https://query.sse.com.cn/sseQuery/commonQuery.do",
                     map[string]string{
                                "Host": "query.sse.com.cn",
                                "Pragma": "no-cache",
                                "Referer": "https://www.sse.com.cn/assortment/stock/list/share/",
                                "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36",
                    },
                    map[string]string{
                         "STOCK_TYPE": "1",
                         "REG_PROVINCE": "",
                         "CSRC_CODE": "",
                         "STOCK_CODE": "",
                         "sqlId": "COMMON_SSE_CP_GPJCTPZ_GPLB_GP_L",
                         "COMPANY_STATUS": "2,4,5,7,8",
                         "type": "inParams",
                         "isPagination": "",
                         "pageHelp.cacheSize": "1",
                         "pageHelp.beginPage": "1",
                         "pageHelp.pageSize": "1",
                         "pageHelp.pageNo": "1",
                         "pageHelp.endPage": "1",
                    },
                },
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
                        var respData responseData
			err := Get(tt.Url, tt.Headers, tt.Params, &respData)
			if err != nil {
			    t.Logf("error: %v.\n", err)
			} else {
			    for _, s := range respData.Result {
				t.Logf("%s, %s, %s, %s\n", s.AStockCode, s.CompanyAbbr, s.ListDate, s.CSRCCodeDesc)
   		  	    }
			}
		})
	}
}
