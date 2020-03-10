# 数据库设计文档

## 需求

### cookie 储存

当微信用户打开小程序时，会发来一个code,此时，后端需要做两件事

1. 向微信api发送请求，确定是合法用户，微信会返回用户的openid  此时我们将用户的openid作为用户的唯一标志，也就是数据库表中的主码
2. 生成session分发给用户，然后该session将会作为该用户发送请求时的标志。

这里有两种做法。

1. 将session储存
2. 不存，只校验时间合法性，所以。在middleware的session已经帮我们做了合法校验，无需再次存，这样也减少了不断读取数据库的时间。

### 上报人员信息

用户需要绑定信息，也就是说，帮定信息就是创建一个数据条目的过程,在这里为了简便，不再拆分表去满足第二范式。

``` go
// WeChat means Who has loged in  
type WeChat struct{
  gorm.Model
  OpenID string `gorm:"unique;"`
}
// Reporter 上报人
type Reporter struct {
  gorm.Model
  WeChat  WeChat `gorm:"association_foreignkey:WeChatRefer;unique;type:varchar(200)"`
  WeChatRefer string //对应的微信号
  OrgID  int  // 机构id
  Name string`gorm:"varchar(30);"`
  Gender int  //0代表女性，1代表男性
  Tel int64    //手机号
  StuNum int64 //学号
  LastUpdateTime *time.Time`gorm:"type:datatime;default:null;"` //最后更新时间
}

```

### 上报记录

``` go
// Record 上报记录
type Record struct{
  gorm.Model
  Reporter Reporter  `gorm:"foreign_key:ReporterRefer;`
  ReporterRefer  int
  IsReturnSchool int `gorm:"not null;"` //是否返校（选项）
  CurrentHealthValue int //当前身体状况
  CurrentDistrictPath int  //当前所在地
  CurrentContagionRiskValue int  //传染风险
  PsyStatus int //心理状况
  PsyDemand int  //心理需求
  PsyKnowledge int  //所需心理知识
  CurrentTemperature int //今日体温
}
```
