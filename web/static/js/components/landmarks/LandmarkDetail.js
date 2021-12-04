import React from 'react';
import {useDispatch, useSelector} from 'react-redux';
import {deleteLmForm, editForm} from '../../store/formsSlice';
import {closeLmDetail} from '../../store/landmarkSlice';

const LandmarkDetail = () => {
  const dispatch = useDispatch();
  const {model, onEdit, onDeleteLm} = useSelector((state) => state.forms);
  const {
    selectedLmStatus,
    selectedLm,
  } = useSelector((state) => state.landmarks);

  const formEditHandle = () => {
    if (!onEdit || model !='LANDMARK') {
      dispatch(editForm('LANDMARK'));
    }
  };

  const formDeleteHandle = () => {
    if (!onDeleteLm || model != 'DELETE') {
      dispatch(deleteLmForm('DELETE'));
    }
  };

  const closeLmDetailHandle = () => {
    if (selectedLmStatus) {
      dispatch(closeLmDetail());
    }
  };

  const detail = () => (
    <article className="landmark-detail">
      <div className="landmark-img">
        <button
          type="button"
          onClick={closeLmDetailHandle}
        >
            X
        </button>
        <img src={selectedLm.imgURL} alt={selectedLm.name} />
      </div>
      <div className="landmark-description">
        <h2>{selectedLm.name}</h2>
        <h3>{selectedLm.nativeName}</h3>
        <p>{selectedLm.country}</p>
        <p>{selectedLm.stateCity}</p>
        <div className="wiki">
          <h4>{selectedLm.type}</h4>
          <a href={selectedLm.wikiURL}><button>Wikipedia</button></a>
        </div>
        <p>{selectedLm.description}</p>
      </div>
      <div className="detail-con">
        <table>
          <thead>
            <tr>
              <th>Start Year</th>
              <th>End Year</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>{selectedLm.startYear}</td>
              <td>{selectedLm.endYear}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div className="detail-con">
        <table>
          <thead>
            <tr>
              <th>Area</th>
              <th>Height</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>{selectedLm.area}</td>
              <td>{selectedLm.height}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div className="landmark-actions">
        <button type="button" onClick={formEditHandle}>Edit</button>
        <button type="button" onClick={formDeleteHandle}>Delete</button>
      </div>
    </article>
  );
  return (
    <div>
      {selectedLmStatus && detail()}
    </div>
  );
};

export default LandmarkDetail;
