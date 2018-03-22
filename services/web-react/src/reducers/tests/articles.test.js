import articleReducer from '../articles';

describe('Reducers', () => {
  describe('Articles', () => {

    describe('Default', () => {
      it('should return the default state', () => {
        const action = {
          type: 'something else'
        };
        const state = articleReducer(undefined, action);
        expect(state).toEqual({
          index: {
            articles: []
          },
          show: {
            resource: null
          }
        });

      });
    })
  });
});
