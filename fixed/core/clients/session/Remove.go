package Sessions

func (s *Session) Remove() {
	SessionMutex.Lock()
	delete(Sessions, s.ID)
	SessionMutex.Unlock()
}