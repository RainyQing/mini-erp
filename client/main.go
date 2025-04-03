package main

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//go:embed ttf/msyh.ttf
var fontData []byte // 字体文件内容将嵌入到二进制中

// 自定义主题结构体
type myTheme struct {
	defaultTheme fyne.Theme // 嵌入默认主题（实现其他方法）
}

// 实现 theme.Theme 接口的 Font 方法
func (m *myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return &fyne.StaticResource{
		StaticName:    "msyh.ttf",
		StaticContent: fontData,
	}
}

// 其他必要方法（直接委托给默认主题）
func (m *myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return m.defaultTheme.Color(name, variant)
}

func (m *myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return m.defaultTheme.Icon(name)
}

func (m *myTheme) Size(name fyne.ThemeSizeName) float32 {
	return m.defaultTheme.Size(name)
}

func main() {
	a := app.New()
	customTheme := &myTheme{defaultTheme: theme.DefaultTheme()}
	a.Settings().SetTheme(customTheme)

	w := a.NewWindow("Mini ERP System")
	w.Resize(fyne.NewSize(1024, 768))

	// 预加载各模块内容
	inventoryContent := createInventoryTab()
	orderContent := createOrderTab()
	financeContent := createFinanceTab()

	// 创建可更新的内容容器
	// 初始化主内容容器
	mainContent := container.NewMax(inventoryContent)

	// 创建导航按钮（使用正确的容器引用）
	navButtons := container.NewVBox(
		widget.NewButton("库存管理", func() { mainContent.Objects = []fyne.CanvasObject{inventoryContent} }),
		widget.NewButton("订单管理", func() { mainContent.Objects = []fyne.CanvasObject{orderContent} }),
		widget.NewButton("财务管理", func() { mainContent.Objects = []fyne.CanvasObject{financeContent} }),
	)

	// 构建最终布局
	content := container.NewBorder(nil, nil, navButtons, nil, mainContent)
	w.SetContent(content)

	w.ShowAndRun()
}

func createInventoryTab() *fyne.Container {
	// 创建产品列表
	productList := widget.NewList(
		func() int { return 0 }, // 临时返回0，后续连接后端API获取实际数据
		func() fyne.CanvasObject {
			return widget.NewLabel("Product Item")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			// 后续实现产品数据绑定
		},
	)

	// 创建添加产品按钮
	addButton := widget.NewButton("添加产品", func() {
		// 后续实现添加产品功能
	})

	return container.NewBorder(nil, addButton, nil, nil, productList)
}

func createOrderTab() *fyne.Container {
	// 创建订单列表
	orderList := widget.NewList(
		func() int { return 0 }, // 临时返回0，后续连接后端API获取实际数据
		func() fyne.CanvasObject {
			return widget.NewLabel("Order Item")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			// 后续实现订单数据绑定
		},
	)

	// 创建新建订单按钮
	newOrderButton := widget.NewButton("新建订单", func() {
		// 后续实现新建订单功能
	})

	return container.NewBorder(nil, newOrderButton, nil, nil, orderList)
}

func createFinanceTab() *fyne.Container {
	// 创建财务记录列表
	financeList := widget.NewList(
		func() int { return 0 }, // 临时返回0，后续连接后端API获取实际数据
		func() fyne.CanvasObject {
			return widget.NewLabel("Finance Record")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			// 后续实现财务数据绑定
		},
	)

	// 创建添加记录按钮
	addRecordButton := widget.NewButton("添加记录", func() {
		// 后续实现添加财务记录功能
	})

	return container.NewBorder(nil, addRecordButton, nil, nil, financeList)
}
