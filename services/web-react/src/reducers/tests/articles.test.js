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
                        pagination: {
                            currentPage: 0,
                            totalPages: 0
                        },
                        resources: [],
                        selected: {},
                        archives: {},
                        tags: []
                    },
                    show: {
                        resource: null
                    }
                });

            });
        })
    });
});
