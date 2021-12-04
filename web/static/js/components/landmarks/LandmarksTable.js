import React, {useEffect} from 'react';
import {useDispatch, useSelector} from 'react-redux';
import {fetchLandmarks} from '../../store/landmarkSlice';
import LandmarkRow from './LandmarkRow';

const LandmarkPanel = () => {
  const dispatch = useDispatch();
  const {
    landmarksStatus,
    landmarksArr,
  } = useSelector((state) => state.landmarks);

  useEffect(() => {
    if (!landmarksStatus) {
      setTimeout(() => dispatch(fetchLandmarks()), 300);
    }
  }, []);

  const mapLandmarks = landmarksArr.map((lm) => (
    <LandmarkRow key={lm.id} landmark={lm} />
  ));

  return (
    <div className="landmarks-con">
      <table>
        <thead>
          <tr>
            <th>+</th>
            <th>ID</th>
            <th>Name</th>
            <th>Native Name</th>
            <th>Type</th>
            <th>Continent</th>
            <th>Country</th>
            <th>City</th>
            <th>Creator</th>
          </tr>
        </thead>
        <tbody>
          {mapLandmarks}
        </tbody>
      </table>
    </div>
  );
};

export default LandmarkPanel;
