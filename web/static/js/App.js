import React from 'react';
import UserPanel from './components/users/UserPanel';
import LandmarkPanel from './components/landmarks/LandmarksTable';
import LandmarkDetail from './components/landmarks/LandmarkDetail';

const App = () => (
  <div>
    <UserPanel />
    <main>
      <div>
        <LandmarkPanel />
        <LandmarkDetail />
      </div>
    </main>
  </div>
);

export default App;
