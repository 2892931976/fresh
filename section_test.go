package main

// func TestNewSection(t *testing.T) {
// 	var s *section

// 	s = newSection("stylesheets: *.css, *.less, , , ")
// 	assert.Equal(t, "stylesheets", s.Name)
// 	assert.Equal(t, []string{"*.css", "*.less"}, s.Globs)
// 	assert.Equal(t, 0, len(s.Commands))

// 	// only name, without globs
// 	s = newSection("foo-section")
// 	assert.Equal(t, "foo-section", s.Name)
// 	assert.Equal(t, 0, len(s.Globs))
// 	assert.Equal(t, 0, len(s.Commands))
// }

// func TestSection_NewCommand(t *testing.T) {
// 	s := newSection("go")
// 	assert.Equal(t, 0, len(s.Commands))
// 	c := s.NewCommand("build", "./build")
// 	assert.Equal(t, 1, len(s.Commands))
// 	assert.Equal(t, c, s.Commands[0])
// }
