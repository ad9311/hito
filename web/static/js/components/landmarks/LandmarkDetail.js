import React from 'react';
import {useDispatch, useSelector} from 'react-redux';
import {deleteLmForm, editForm} from '../../store/formsSlice';

const LandmarkDetail = () => {
  const dispatch = useDispatch();
  const {type, onEdit, onDeleteLm} = useSelector((state) => state.forms);
  const {
    selectedLmStatus,
    selectedLm,
  } = useSelector((state) => state.landmarks);

  const formEditHandle = () => {
    if (!onEdit || type !='LANDMARK') {
      dispatch(editForm('LANDMARK'));
    }
  };

  const formDeleteHandle = () => {
    if (!onDeleteLm || type != 'DELETE') {
      dispatch(deleteLmForm('DELETE'));
    }
  };

  const detail = () => (
    <article className="landmark-detail">
      <div className="landmark-img">
        <img src={selectedLm.imgURL} alt={selectedLm.name} />
      </div>
      <div className="landmark-description">
        <h2>{selectedLm.name}</h2>
        <h3>{selectedLm.nativeName}</h3>
        <div className="wiki">
          <h4>{selectedLm.type}</h4>
          <a href={selectedLm.wikiURL}><button>Wikipedia</button></a>
        </div>
        <p>{selectedLm.description}</p>
      </div>
      <div className="detail-con">
        <p>{selectedLm.country}</p>
        <p>{selectedLm.city}</p>
      </div>
      <div className="detail-con">
        <table>
          <thead>
            <tr>
              <th>Latitude</th>
              <th>Longitude</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>{selectedLm.latitude}</td>
              <td>{selectedLm.longitude}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div className="detail-con">
        <table>
          <thead>
            <tr>
              <th>Length</th>
              <th>Width</th>
              <th>Height</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>{selectedLm.length}</td>
              <td>{selectedLm.width}</td>
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
