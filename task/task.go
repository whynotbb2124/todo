package task

var Tasks = make(map[string]map[string]Task)

type Task struct {
  Title       string
  Description string
  Status      bool
  Deadline    string
  UserLogin   string
}

func Create(title, description, deadline, userLogin string) {
  var t Task
  t.Title = title
  t.Description = description
  t.Status = false
  t.Deadline = deadline
  t.UserLogin = userLogin

  if Tasks[userLogin] == nil {
    Tasks[userLogin] = make(map[string]Task)
  }

  Tasks[userLogin][title] = t
}

func (t *Task) Update() {
  if Tasks[t.UserLogin] == nil {
    return
  }

  Tasks[t.UserLogin][t.Title] = *t
}

func (t *Task) Delete() {
  if Tasks[t.UserLogin] == nil {
    return
  }

  delete(Tasks[t.UserLogin], t.Title)
}

func Read(userLogin, title string) Task {
  if Tasks[userLogin] == nil {
    return Task{}
  }

  return Tasks[userLogin][title]
}