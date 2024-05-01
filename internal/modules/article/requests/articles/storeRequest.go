package articles

// Binding from html form
type StoreRequest struct {
	Title   string `form:"title" binding:"required,min=3,max=50"`
	Content string `form:"content"  binding:"required,min=10,max=500"`
}
