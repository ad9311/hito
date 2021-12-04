import React from 'react';
import UserPanel from './components/users/UserPanel';
import LandmarkPanel from './components/landmarks/LandmarksTable';
import LandmarkDetail from './components/landmarks/LandmarkDetail';
import Forms from './components/forms/Forms';

const App = () => (
  <div>
    <UserPanel />
    <main className="main-home">
      <LandmarkPanel />
      <Forms />
    </main>
    <LandmarkDetail />
  </div>
);

export default App;
