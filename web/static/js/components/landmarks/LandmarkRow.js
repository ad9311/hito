import React from 'react';
import PropTypes from 'prop-types';
import {useDispatch, useSelector} from 'react-redux';
import {selectLm} from '../../store/landmarkSlice';

const LandmarkRow = (props) => {
  const dispatch = useDispatch();
  const {
    selectedLm,
    selectedLmStatus,
  } = useSelector((state) => state.landmarks);
  const {landmark} = props;

  const selectLandmark = () => {
    if (selectedLm.name !== landmark.name || !selectedLmStatus) {
      dispatch(selectLm(landmark));
    }
  };

  return (
    <tr>
      <td>
        <button type="button" onClick={selectLandmark}>
          +
        </button>
      </td>
      <td>{landmark.id}</td>
      <td>{landmark.name}</td>
      <td>{landmark.nativeName}</td>
      <td>{landmark.type}</td>
      <td>{landmark.continent}</td>
      <td>{landmark.country}</td>
      <td>{landmark.stateCity}</td>
      <td>{landmark.userUsername}</td>
    </tr>
  );
};

LandmarkRow.propTypes = {
  landmark: PropTypes.objectOf(PropTypes.objectOf),
};

LandmarkRow.defaultProps = {
  landmark: {},
};

export default LandmarkRow;
